package cryptoarticle_business

import (
	"mini-news/app/global/errorcode"
	"mini-news/app/global/helper"
	"mini-news/app/service/cryptonews_service"
)

func (b *business) SyncAbmediaArticles() (goErr errorcode.Error) {
	insertFields := []map[string]interface{}{}

	abArticles, goErr := b.cryptoNewsService.GetAbmediaArticles()
	if goErr != nil {
		return
	}

	names := filterNames(abArticles)

	abData, goErr := b.cryptoArticleRepo.GetArticleByTypeNames(cryptonews_service.ABMEDIA, names)
	if goErr != nil {
		return
	}

	if len(abData) == 0 {
		// 爬出來的資料DB都沒有 全部寫入
		for k := range abArticles {
			myMap, err := helper.StructToMap(abArticles[k])

			if err != nil {
				goErr = helper.ErrorHandle(errorcode.ErrorBusiness, errorcode.StructToMapError, err.Error())
				return
			}
			insertFields = append(insertFields, myMap)
		}
	} else {
		for k := range abArticles {
			for i := range abData {
				if abArticles[k].Name == abData[i].Name {
					continue
				}

				myMap, err := helper.StructToMap(abArticles[k])
				if err != nil {
					goErr = helper.ErrorHandle(errorcode.ErrorBusiness, errorcode.StructToMapError, err.Error())
					return
				}

				insertFields = append(insertFields, myMap)
			}
		}
	}

	goErr = b.cryptoArticleRepo.InsertArticleByMaps(insertFields)
	if goErr != nil {
		return
	}

	return
}

func (b *business) SyncBlockTempoArticles() (goErr errorcode.Error) {
	insertFields := []map[string]interface{}{}

	btArticles, goErr := b.cryptoNewsService.GetBlockTempoArticles()
	if goErr != nil {
		return
	}

	names := filterNames(btArticles)

	btData, goErr := b.cryptoArticleRepo.GetArticleByTypeNames(cryptonews_service.BLOCKTEMPO, names)
	if goErr != nil {
		return
	}

	if len(btData) == 0 {
		for k := range btArticles {
			myMap, err := helper.StructToMap(btArticles[k])
			if err != nil {
				goErr = helper.ErrorHandle(errorcode.ErrorBusiness, errorcode.StructToMapError, err.Error())
				return
			}
			insertFields = append(insertFields, myMap)
		}
	} else {
		for k := range btArticles {
			for i := range btData {
				if btArticles[k].Name == btData[i].Name {
					continue
				}

				myMap, err := helper.StructToMap(btArticles[k])
				if err != nil {
					goErr = helper.ErrorHandle(errorcode.ErrorBusiness, errorcode.StructToMapError, err.Error())
					return
				}

				insertFields = append(insertFields, myMap)
			}
		}
	}

	goErr = b.cryptoArticleRepo.InsertArticleByMaps(insertFields)
	if goErr != nil {
		return
	}

	return
}

func (b *business) SyncBlockCastArticles() (goErr errorcode.Error) {
	insertFields := []map[string]interface{}{}

	bcArticles, goErr := b.cryptoNewsService.GetBlockCastArticles()
	if goErr != nil {
		return
	}

	names := filterNames(bcArticles)

	bcData, goErr := b.cryptoArticleRepo.GetArticleByTypeNames(cryptonews_service.BLOCKTEMPO, names)
	if goErr != nil {
		return
	}

	if len(bcData) == 0 {
		for k := range bcArticles {
			myMap, err := helper.StructToMap(bcArticles[k])
			if err != nil {
				goErr = helper.ErrorHandle(errorcode.ErrorBusiness, errorcode.StructToMapError, err.Error())
				return
			}
			insertFields = append(insertFields, myMap)
		}
	} else {
		for k := range bcArticles {
			for i := range bcData {
				if bcArticles[k].Name == bcData[i].Name {
					continue
				}

				myMap, err := helper.StructToMap(bcArticles[k])
				if err != nil {
					goErr = helper.ErrorHandle(errorcode.ErrorBusiness, errorcode.StructToMapError, err.Error())
					return
				}

				insertFields = append(insertFields, myMap)
			}
		}
	}

	goErr = b.cryptoArticleRepo.InsertArticleByMaps(insertFields)
	if goErr != nil {
		return
	}

	return
}

func filterNames(articles []cryptonews_service.CryptoArticle) (names []string) {
	for k := range articles {
		names = append(names, articles[k].Name)
	}
	return
}
