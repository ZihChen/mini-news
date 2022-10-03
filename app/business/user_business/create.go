package user_business

import (
	"mini-news/app/global/errorcode"
	"mini-news/app/global/helper"
	"mini-news/app/global/structs/request"
)

func (b *business) CreateUser(option request.RegisterUserOption) (goErr errorcode.Error) {
	// 檢查username是否存在
	exist, goErr := b.userRepo.CheckUserExistByUserName(option.Username)
	if goErr != nil {
		return
	}

	if exist {
		goErr = helper.ErrorHandle(errorcode.ErrorBusiness, errorcode.UserAlreadyExist, "")
		return
	}

	hashPwd, err := helper.HashPassword(option.Password)
	if err != nil {
		goErr = helper.ErrorHandle(errorcode.ErrorBusiness, errorcode.HashPasswordError, err.Error())
		return
	}

	option.Password = hashPwd
	mapFields, err := helper.StructToMap(option)
	if err != nil {
		goErr = helper.ErrorHandle(errorcode.ErrorBusiness, errorcode.StructToMapError, err.Error())
		return
	}

	delete(mapFields, "password_confirmation")

	if goErr = b.userRepo.CreateUserByMap(mapFields); goErr != nil {
		return
	}

	return
}
