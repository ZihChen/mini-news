package article_tag_repo

import (
	"mini-news/app/global/errorcode"
	"mini-news/app/global/helper"
	"mini-news/app/model"
)

func (r *repo) CreateArticleTagByMap(insertFields []map[string]interface{}) (goErr errorcode.Error) {
	db, goErr := r.DB.DBConnect()
	if goErr != nil {
		return
	}

	tx := db.Begin()

	if err := tx.Model(&model.ArticleTag{}).Create(insertFields).Error; err != nil {
		tx.Rollback()
		goErr = helper.ErrorHandle(errorcode.ErrorRepository, errorcode.ArticleTagCreateError, err.Error())
		return
	}

	tx.Commit()

	return
}
