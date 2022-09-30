package cryptoarticle_repo

import (
	"mini-news/app/global/errorcode"
	"mini-news/app/global/helper"
	"mini-news/app/model"
)

func (r *repo) GetArticleByTypeNames(site string, names []string) (articles []model.CryptoArticle, goErr errorcode.Error) {
	db, goErr := r.DB.DBConnect()
	if goErr != nil {
		return
	}

	if err := db.Where("site = ? AND name IN ?", site, names).Find(&articles).Error; err != nil {
		goErr = helper.ErrorHandle(errorcode.ErrorRepository, errorcode.GetArticleByTypeNamesError, err.Error())
		return
	}

	return
}
