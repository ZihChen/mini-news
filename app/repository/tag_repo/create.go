package tag_repo

import (
	"mini-news/app/global/errorcode"
	"mini-news/app/global/helper"
	"mini-news/app/model"
)

func (r *repo) CreateTagByMap(fields map[string]interface{}) (goErr errorcode.Error) {
	db, goErr := r.DB.DBConnect()
	if goErr != nil {
		return
	}

	tx := db.Begin()

	if err := tx.Model(&model.Tag{}).Create(fields).Error; err != nil {
		goErr = helper.ErrorHandle(errorcode.ErrorRepository, errorcode.TagCreateError, err.Error())
		tx.Rollback()
		return
	}

	tx.Commit()

	return
}
