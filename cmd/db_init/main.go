package main

import (
	"fmt"
	"log"
	"github.com/wuchuwuyou/go-web-demo/model"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main()  {
	log.Println("DB init...")
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)

	db.DropTableIfExists(model.User{},model.Post{})
	db.CreateTable(model.User{},model.Post{})

	users := []model.User {
		{
			Username : "murphy",
			PasswordHash: model.GeneratePasswordHash("abc123"),
			Email: "murphy@murphy.com",
			Avatar: fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=identicon", model.Md5("murphy@murphy.com")),
			Posts: []model.Post{
				{Body:"Have a nice day!"},
			},
		},
		{
			Username:     "rene",
            PasswordHash: model.GeneratePasswordHash("abc123"),
			Email:        "rene@test.com",
			Avatar:       fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=identicon", model.Md5("rene@test.com")),
            Posts: []model.Post{
                {Body: "The Avengers movie was so cool!"},
                {Body: "Sun shine is beautiful"},
            },
		},
	}
	
	for _, u := range users {
		db.Debug().Create(&u)
	}
}