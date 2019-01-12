package main
import (
	"net/http"
	"github.com/wuchuwuyou/go-web-demo/controller"
	"github.com/wuchuwuyou/go-web-demo/model"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func main()  {

	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)
	
	controller.Startup()
	http.ListenAndServe(":8888",nil)
}