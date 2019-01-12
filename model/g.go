package model

import (
	"log"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/wuchuwuyou/go-web-demo/config"
)

var db *gorm.DB

func SetDB(database *gorm.DB)  {
	db = database
}

func ConnectToDB() *gorm.DB {
	connectingStr := config.GetMysqlConnectingString()
	fmt.Println(connectingStr)
	log.Println("Connet to DB...")
	db,err := gorm.Open("mysql",connectingStr)
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect database")
	}
	db.SingularTable(true)
	return db
}