package article_tag_business

import (
	"mini-news/app/global/errorcode"
	"mini-news/app/global/structs/request"
	"mini-news/app/repository/article_tag_repo"
	"mini-news/app/repository/tag_repo"
	"sync"
)

type Interface interface {
	CreateArticleTag(option request.CreateArticleTagOption) (goErr errorcode.Error)
}

type business struct {
	articleTagRepo article_tag_repo.Interface
	tagRepo        tag_repo.Interface
}

var singleton *business
var once sync.Once

func NewBusiness() Interface {
	once.Do(func() {
		singleton = &business{
			articleTagRepo: article_tag_repo.NewRepo(),
			tagRepo:        tag_repo.NewRepo(),
		}
	})
	return singleton
}
