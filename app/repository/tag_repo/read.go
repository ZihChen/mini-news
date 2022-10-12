package tag_repo

import (
	"mini-news/app/global/errorcode"
	"mini-news/app/global/helper"
	"mini-news/app/model"
)

func (r *repo) GetTagsCountByTagIDsAndUserID(tagIDs []int, userID int) (count int64, goErr errorcode.Error) {
	db, goErr := r.DB.DBConnect()
	if goErr != nil {
		return
	}

	if err := db.Model(&model.Tag{}).Where("id IN ? AND user_id = ?", tagIDs, userID).Count(&count).Error; err != nil {
		goErr = helper.ErrorHandle(errorcode.ErrorRepository, errorcode.GetTagsCountError, err.Error())
		return
	}

	return
}
