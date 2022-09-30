package cryptoarticle_repo

import (
	"mini-news/app/global/errorcode"
	"mini-news/app/global/helper"
	"mini-news/app/model"
)

func (r *repo) InsertArticleByMaps(maps []map[string]interface{}) (goErr errorcode.Error) {
	db, goErr := r.DB.DBConnect()
	if goErr != nil {
		return
	}

	tx := db.Begin()

	if err := tx.Model(&model.CryptoArticle{}).Create(maps).Error; err != nil {
		tx.Rollback()
		goErr = helper.ErrorHandle(errorcode.ErrorRepository, errorcode.InsertArticleByMapsError, err.Error())
		return
	}

	tx.Commit()

	return
}
