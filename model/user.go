package model

import (
	"fmt"
    "time"
    "log"
)

// User struct
type User struct {
    ID              int  `gorm:"primary_key"`
    Username        string `gorm:"type:varchar(64)"`
    Email           string `gorm:"type:varchar(120)"`
	PasswordHash    string `gorm:"type:varchar(128)"`
    Posts           []Post
    Followers       []*User `gorm:"many2many:follower;association_jointable_foreignkey:follower_id"`
    LastSeen        *time.Time 
    AboutMe         string `gorm:"type:varchar(140)"`
    Avatar          string `gorm:"type:varchar(200)"`
}

func (u *User) SetPassword(password string)  {
    u.PasswordHash = GeneratePasswordHash(password)
}

func (u *User) CheckPassword(password string) bool {
    return GeneratePasswordHash(password) == u.PasswordHash
}

func GetUserByUsername(username string) (*User,error) {
    var user User
    if err := db.Where("username=?",username).Find(&user).Error; err != nil {
        return nil,err
    }
    return &user,nil
}

func (u *User) SetAvatar(email string) {
    u.Avatar = fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=identicon", Md5(email))
}

func AddUser(username,password,email string) error {
    user := User{Username: username, Email: email}
    user.SetPassword(password)
    user.SetAvatar(email)
    if err := db.Create(&user).Error;err != nil {
        return err
    }
    return user.FollowSelf() 
}

func UpdateUserByUsername(username string,contents map[string]interface{}) error {
    item,err := GetUserByUsername(username)
    if err != nil {
        return err
    }
    return db.Model(item).Updates(contents).Error
}

func UpdateLastSeen(username string) error {
    contents := map[string]interface{}{"last_seen":time.Now()}
    return UpdateUserByUsername(username,contents)
}

func UpdateAboutMe(username, text string) error {
    contents := map[string]interface{}{"about_me":text}
    return UpdateUserByUsername(username,contents)
}

func (u *User) Follow(username string) error {
    other,err := GetUserByUsername(username)
    if err != nil {
        return err
    }
    return db.Model(other).Association("Followers").Append(u).Error
}

func (u *User) Unfollow(username string) error {
    other, err := GetUserByUsername(username)
	if err != nil {
		return err
	}
	return db.Model(other).Association("Followers").Delete(u).Error
}

// FollowSelf func
func (u *User) FollowSelf() error {
	return db.Model(u).Association("Followers").Append(u).Error
}

// FollowersCount func
func (u *User) FollowersCount() int {
	return db.Model(u).Association("Followers").Count()
}

// FollowingIDs func
func (u *User) FollowingIDs() []int {
	var ids []int
	rows, err := db.Table("follower").Where("follower_id = ?", u.ID).Select("user_id, follower_id").Rows()
	if err != nil {
		log.Println("Counting Following error:", err)
		return ids
	}
	defer rows.Close()
	for rows.Next() {
		var id, followerID int
		rows.Scan(&id, &followerID)
		ids = append(ids, id)
	}
	return ids
}

// FollowingCount func
func (u *User) FollowingCount() int {
	ids := u.FollowingIDs()
	return len(ids)
}

// FollowingPosts func
func (u *User) FollowingPosts() (*[]Post, error) {
	var posts []Post
	ids := u.FollowingIDs()
	if err := db.Preload("User").Order("timestamp desc").Where("user_id in (?)", ids).Find(&posts).Error; err != nil {
		return nil, err
	}
	return &posts, nil
}

// IsFollowedByUser func
func (u *User) IsFollowedByUser(username string) bool {
	user, _ := GetUserByUsername(username)
	ids := user.FollowingIDs()
	for _, id := range ids {
		if u.ID == id {
			return true
		}
	}
	return false
}

// CreatePost func
func (u *User) CreatePost(body string) error {
	post := Post{Body: body, UserID: u.ID}
	return db.Create(&post).Error
}

func (u *User) FollowingPostsByPageAndLimit(page,limit int)(*[]Post,int,error) {
    var total int 
    var posts []Post
    offset := (page - 1) * limit
    ids := u.FollowingIDs()
    if err := db.Preload("User").Order("timestamp desc").Where("user_id in (?)",ids).Offset(offset).Limit(limit).Find(&posts).Error;err != nil {
        return nil,total,err
    }
    db.Model(&Post{}).Where("user_id in (?)",ids).Count(&total)
    return &posts,total,nil
}