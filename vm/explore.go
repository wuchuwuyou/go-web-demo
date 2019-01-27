package	vm

import (
	"github.com/wuchuwuyou/go-web-demo/model"
)

type ExploreViewModel struct {
	BaseViewModel
	Posts []model.Post
	BasePageViewModel
}

type ExploreViewModelOp struct {}

func (ExploreViewModelOp) GetVM(username string,page,limit int) ExploreViewModel {
	posts,total,_ := model.GetPostsByPageAndLimit(page,limit)
	v := ExploreViewModel{}
	v.SetTitle("Explore")
	v.Posts = *posts
	v.SetBasePageViewModel(total,page,limit)
	v.SetCurrentUser(username)
	return v
}