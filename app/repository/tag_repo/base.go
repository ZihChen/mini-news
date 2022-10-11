package tag_repo

import (
	"mini-news/app/global/errorcode"
	"mini-news/internal/database"
	"sync"
)

type Interface interface {
	CreateTagByMap(fields map[string]interface{}) (goErr errorcode.Error)
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
