package vm
import (
	"log"
	"github.com/wuchuwuyou/go-web-demo/model"
)
type ResetPasswordRequestViewModel struct {
	LoginViewModel
}

type ResetPasswordRequestViewModelOp struct {}

func (ResetPasswordRequestViewModelOp) GetVM() ResetPasswordRequestViewModel {
	v := ResetPasswordRequestViewModel{}
	v.SetTitle("Forget Password")
	return v
}

func CheckEmailExist(email string) bool {
	_,err := model.GetUserByEmail(email)
	if err != nil {
		log.Println("Can not find eamil:",email)
		return false
	}
	return true
}