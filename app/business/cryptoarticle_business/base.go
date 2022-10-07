package cryptoarticle_business

import (
	"mini-news/app/global/errorcode"
	"mini-news/app/repository/cryptoarticle_repo"
	"mini-news/app/service/cryptonews_service"
	"sync"
)

type Interface interface {
	// SyncAbmediaArticles 同步鏈新聞資料
	SyncAbmediaArticles() (goErr errorcode.Error)
	// SyncBlockTempoArticles 同步動區新聞資料
	SyncBlockTempoArticles() (goErr errorcode.Error)
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
