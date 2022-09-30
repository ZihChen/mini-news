package cryptoarticle_repo

import (
	"mini-news/app/global/errorcode"
	"mini-news/app/model"
	"mini-news/internal/database"
	"sync"
)

type Interface interface {
	GetArticleByTypeNames(site string, names []string) (articles []model.CryptoArticle, goErr errorcode.Error)
	InsertArticleByMaps(maps []map[string]interface{}) (goErr errorcode.Error)
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
