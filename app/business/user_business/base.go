package user_business

import (
	"mini-news/app/global/errorcode"
	"mini-news/app/global/structs/request"
	"mini-news/app/repository/user_repo"
	"sync"
)

type Interface interface {
	CreateUser(option request.CreateUserOption) (goErr errorcode.Error)
}

type business struct {
	userRepo user_repo.Interface
}

var singleton *business
var once sync.Once

func NewBusiness() Interface {
	once.Do(func() {
		singleton = &business{
			userRepo: user_repo.NewRepo(),
		}
	})
	return singleton
}
