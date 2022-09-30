package cryptonews_service

import (
	"mini-news/app/global/errorcode"
	"sync"
)

type Interface interface {
	// GetAbmediaArticles 爬取鏈新聞
	GetAbmediaArticles() (articles []CryptoArticle, goErr errorcode.Error)
	// GetBlockTempoArticles 爬取動區新聞
	GetBlockTempoArticles() (articles []CryptoArticle, goErr errorcode.Error)
	// GetBlockCastArticles 區塊客
	GetBlockCastArticles() (articles []CryptoArticle, goErr errorcode.Error)
}

type service struct{}

type CryptoArticle struct {
	Name   string `json:"name"`
	Site   string `json:"site"`
	Image  string `json:"image"`
	Title  string `json:"title"`
	Source string `json:"source"`
	Date   string `json:"date"`
}

var singleton *service
var once sync.Once

func NewService() Interface {
	once.Do(func() {
		singleton = &service{}
	})
	return singleton
}
