package user_repo

import (
	"mini-news/app/global/errorcode"
	"mini-news/app/global/helper"
	"mini-news/app/model"
)

func (r *repo) CreateUserByMap(fields map[string]interface{}) (goErr errorcode.Error) {
	db, goErr := r.DB.DBConnect()
	if goErr != nil {
		return
	}

	tx := db.Begin()

	if err := tx.Model(&model.User{}).Create(fields).Error; err != nil {
		goErr = helper.ErrorHandle(errorcode.ErrorRepository, errorcode.UserCreateError, err.Error())
		tx.Rollback()
		return
	}

	tx.Commit()

	return
}
