package user_business

import (
	"mini-news/app/global/helper"
	"mini-news/app/global/structs/request"
)

func (b *business) CreateUser(option request.CreateUserOption) (err error) {

	mapFields, _ := helper.StructToMap(option)

	b.userRepo.CreateUserByMap(mapFields)

	return
}
