package user_repo

import (
	"mini-news/internal/database"
	"sync"
)

type Interface interface {
	CreateUserByMap(fields map[string]interface{}) (err error)
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
