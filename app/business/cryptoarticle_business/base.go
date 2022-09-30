package cryptoarticle_business

import (
	"mini-news/app/global/errorcode"
	"mini-news/app/repository/cryptoarticle_repo"
	"mini-news/app/service/cryptonews_service"
	"sync"
)

type Interface interface {
	SyncAbmediaArticles() (goErr errorcode.Error)
}

type business struct {
	cryptoNewsService cryptonews_service.Interface
	cryptoArticleRepo cryptoarticle_repo.Interface
}

var singleton *business
var once sync.Once

func NewBusiness() Interface {
	once.Do(func() {
		singleton = &business{
			cryptoNewsService: cryptonews_service.NewService(),
			cryptoArticleRepo: cryptoarticle_repo.NewRepo(),
		}
	})
	return singleton
}
