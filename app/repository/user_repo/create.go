package user_repo

import "mini-news/app/model"

func (r *repo) CreateUserByMap(fields map[string]interface{}) (err error) {
	db, err := r.DB.DBConnect()
	if err != nil {
		return
	}

	tx := db.Begin()

	if err = tx.Model(&model.User{}).Create(fields).Error; err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	return
}
