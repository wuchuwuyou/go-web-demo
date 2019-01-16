package vm

import (
	"log"
	"github.com/wuchuwuyou/go-web-demo/model"
)
// LoginViewModel struct
type LoginViewModel struct {
	BaseViewModel
	Errs []string
}

// LoginViewModelOp strutc
type LoginViewModelOp struct{}

// GetVM func
func (LoginViewModelOp) GetVM() LoginViewModel {
    v := LoginViewModel{}
    v.SetTitle("Login")
    return v
}

func (v *LoginViewModel) AddError(errs ...string) {
	v.Errs = append(v.Errs, errs...)
}

func CheckLogin(username,password string) bool {
	user, err := model.GetUserByUsername(username)
	if err != nil {
		log.Println("Can not find username: ",username)
		log.Println("Error:",err)
		return false
	}
	return user.CheckPassword(password)
}