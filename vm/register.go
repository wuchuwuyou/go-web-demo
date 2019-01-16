package vm 

import (
	"log"
	"github.com/wuchuwuyou/go-web-demo/model"
)

type RegisterViewModel struct {
	LoginViewModel
}

type RegisterViewModelOp struct {}

func (RegisterViewModelOp) GetVM() RegisterViewModel {
	v := RegisterViewModel{}
	v.SetTitle("Register")
	return v
}

func CheckUserExist(username string) bool {
	_,err := model.GetUserByUsername(username)
	if err != nil {
		log.Println("Can not find username: ",username)
		return true
	}
	return false
}

func AddUser(username,password,email string) error {
	return model.AddUser(username,password,email)
}