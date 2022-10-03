package user_business

import (
	"mini-news/app/global/errorcode"
	"mini-news/app/global/helper"
	"mini-news/app/global/structs/request"
)

func (b *business) CreateUser(option request.RegisterUserOption) (goErr errorcode.Error) {

	mapFields, err := helper.StructToMap(option)
	if err != nil {
		goErr = helper.ErrorHandle(errorcode.ErrorBusiness, errorcode.StructToMapError, err.Error())
		return
	}

	if goErr = b.userRepo.CreateUserByMap(mapFields); goErr != nil {
		return
	}

	return
}
