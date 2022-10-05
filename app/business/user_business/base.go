package user_business

import (
	"mini-news/app/global/errorcode"
	"mini-news/app/global/structs/request"
	"mini-news/app/repository/user_repo"
	"mini-news/app/service/jwtservice"
	"mini-news/internal/cache"
	"sync"
)

type Interface interface {
	UserLogin(option request.UserLogin) (token string, goErr errorcode.Error)
	CreateUser(option request.RegisterUserOption) (goErr errorcode.Error)
}

type business struct {
	userRepo   user_repo.Interface
	jwtService jwtservice.Interface
	cache      cache.Interface
}

var singleton *business
var once sync.Once

func NewBusiness() Interface {
	once.Do(func() {
		singleton = &business{
			userRepo:   user_repo.NewRepo(),
			jwtService: jwtservice.NewService(),
			cache:      cache.NewRedisConnect(),
		}
	})
	return singleton
}
