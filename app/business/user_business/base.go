package user_business

import (
	"mini-news/app/global/errorcode"
	"mini-news/app/global/structs/request"
	"mini-news/app/repository/user_repo"
	"mini-news/app/service/jwtservice"
	"sync"
)

type Interface interface {
	UserLogin(option request.UserLogin) (token string, goErr errorcode.Error)
	CreateUser(option request.RegisterUserOption) (goErr errorcode.Error)
}

type business struct {
	userRepo   user_repo.Interface
	jwtService jwtservice.Interface
}

var singleton *business
var once sync.Once

func NewBusiness() Interface {
	once.Do(func() {
		singleton = &business{
			userRepo:   user_repo.NewRepo(),
			jwtService: jwtservice.NewService(),
		}
	})
	return singleton
}
