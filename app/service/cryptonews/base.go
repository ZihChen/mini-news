package cryptonews

import "sync"

type Interface interface {
	// GetAbmediaArticles 爬取鏈新聞
	GetAbmediaArticles() (articles []CryptoArticle)
	// GetBlockTempoArticles 爬取動區新聞
	GetBlockTempoArticles() (articles []CryptoArticle)
	// GetBlockCastArticles 區塊客
	GetBlockCastArticles() (articles []CryptoArticle)
}

type service struct{}

type CryptoArticle struct {
	ArticleName string `json:"article_name"`
	Type        string `json:"type"`
	Image       string `json:"image"`
	Title       string `json:"title"`
	Source      string `json:"source"`
	Time        string `json:"time"`
}

var singleton *service
var once sync.Once

func NewService() Interface {
	once.Do(func() {
		singleton = &service{}
	})
	return singleton
}
