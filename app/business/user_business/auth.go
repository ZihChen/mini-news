package user_business

import (
	"mini-news/app/global/consts"
	"mini-news/app/global/errorcode"
	"mini-news/app/global/helper"
	"mini-news/app/global/structs/request"
)

func (b *business) UserLogin(option request.UserLogin) (token string, goErr errorcode.Error) {
	user, goErr := b.userRepo.GetUserByUserName(option.Username)
	if goErr != nil {
		return
	}

	if ok := helper.CheckPasswordHash(user.Password, option.Password); !ok {
		goErr = helper.ErrorHandle(errorcode.ErrorBusiness, errorcode.UserPasswordIncorrectError, "")
		return
	}

	tokenMeta, goErr := b.jwtService.GenerateToken(user)
	if goErr != nil {
		return
	}

	if goErr = b.cache.Set(consts.RedisToken+tokenMeta.Signature, user.Username, consts.TokenExpireTime); goErr != nil {
		return
	}

	return tokenMeta.AccessToken, goErr
}
