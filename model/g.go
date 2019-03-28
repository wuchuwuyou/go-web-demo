package model

import (
	"github.com/jinzhu/gorm"
	"github.com/wuchuwuyou/go-web-demo/config"
	"log"
)

var db *gorm.DB

func SetDB(database *gorm.DB)  {
	db = database
}

func ConnectToDB() *gorm.DB {
	if config.IsHeroku() {
		return ConnectToDBByDBType("postgres", config.GetHerokuConnectingString())
	}
	return ConnectToDBByDBType("mysql", config.GetMysqlConnectingString())
}

func ConnectToDBByDBType(dbtype, connectingStr string) *gorm.DB {
	log.Println("DB Type:", dbtype, "\nConnet to db...")
	db, err := gorm.Open(dbtype, connectingStr)
	if err != nil {
		panic("Failed to connect database")
	}
	db.SingularTable(true)
	return db
}