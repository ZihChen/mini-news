package user_repo

import (
	"mini-news/app/global/errorcode"
	"mini-news/app/model"
	"mini-news/internal/database"
	"sync"
)

type Interface interface {
	CreateUserByMap(fields map[string]interface{}) (goErr errorcode.Error)
	CheckUserExistByUserName(username string) (exist bool, goErr errorcode.Error)
	GetUserByUserName(username string) (user model.User, goErr errorcode.Error)
}

type repo struct {
	DB database.Interface
}

var singleton *repo
var once sync.Once

func NewRepo() Interface {
	once.Do(func() {
		singleton = &repo{
			DB: database.NewDBInstance(),
		}
	})
	return singleton
}
