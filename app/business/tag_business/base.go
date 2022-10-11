package tag_business

import (
	"mini-news/app/global/errorcode"
	"mini-news/app/global/structs/request"
	"mini-news/app/repository/tag_repo"
	"sync"
)

type Interface interface {
	CreateTag(option request.CreateTagOption) (goErr errorcode.Error)
}

type business struct {
	tagRepo tag_repo.Interface
}

var singleton *business
var once sync.Once

func NewBusiness() Interface {
	once.Do(func() {
		singleton = &business{
			tagRepo: tag_repo.NewRepo(),
		}
	})
	return singleton
}
