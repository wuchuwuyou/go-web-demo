package main
import (
	"net/http"
	"github.com/wuchuwuyou/go-web-demo/controller"
)


func main()  {
	controller.Startup()
	http.ListenAndServe(":8888",nil)
}