package model

import (
	"fmt"
    "time"
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
    return db.Create(&user).Error 
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