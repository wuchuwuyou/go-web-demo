package vm

import (
	"log"
	"github.com/wuchuwuyou/go-web-demo/model"
)

// IndexViewModel struct
type IndexViewModel struct {
    BaseViewModel
    model.User
    Posts []model.Post
}

// IndexViewModelOp struct
type IndexViewModelOp struct{}

// GetVM func
func (IndexViewModelOp) GetVM(username string) IndexViewModel {
    u1,_ := model.GetUserByUsername(username)

    posts,err := model.GetPostsByUserID(u1.ID)
    // if err != nil {
        log.Println("get user post error:%s", err)
    //     return 
    // }
    v := IndexViewModel{BaseViewModel{Title: "Homepage"}, *u1, *posts}
    v.SetCurrentUser(username)
    return v
}