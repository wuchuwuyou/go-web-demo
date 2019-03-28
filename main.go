package main

import (
	"github.com/gorilla/context"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/wuchuwuyou/go-web-demo/controller"
	"github.com/wuchuwuyou/go-web-demo/model"
	"log"
	"net/http"
	"os"
)


func main()  {

	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)
	
	controller.Startup()
	//http.ListenAndServe(":8888",context.ClearHandler(http.DefaultServeMux))
	port := os.Getenv("PORT")
	log.Println("Running on port: ", port)
	http.ListenAndServe(":"+port,context.ClearHandler(http.DefaultServeMux))
}
