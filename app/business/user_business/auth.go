package user_business

import (
	"mini-news/app/global/errorcode"
	"mini-news/app/global/helper"
	"mini-news/app/global/structs/request"
)

func (b *business) UserLogin(option request.UserLogin) (token string, goErr errorcode.Error) {
	user, goErr := b.userRepo.GetUserByUserName(option.Username)
	if goErr != nil {
		return
	}

	if user.Password != option.Password {
		goErr = helper.ErrorHandle(errorcode.ErrorBusiness, errorcode.UserPasswordIncorrectError, "")
		return
	}

	token, goErr = b.jwtService.GenerateToken(user)
	if goErr != nil {
		return
	}

	return
}
