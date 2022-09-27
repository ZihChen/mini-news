package user_repo

import (
	"mini-news/app/global/errorcode"
	"mini-news/app/global/helper"
	"mini-news/app/model"
)

func (r *repo) GetUserByUserName(username string) (user model.User, goErr errorcode.Error) {
	db, goErr := r.DB.DBConnect()
	if goErr != nil {
		return
	}

	if err := db.Where("username = ?", username).Find(&user).Error; err != nil {
		goErr = helper.ErrorHandle(errorcode.ErrorRepository, errorcode.GetUserError, err.Error())
		return
	}

	return
}
