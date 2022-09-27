package jwtservice

import (
	"github.com/golang-jwt/jwt"
	"mini-news/app/global/errorcode"
	"mini-news/app/global/helper"
	"mini-news/app/model"
	"time"
)

func (s *service) GenerateToken(user model.User) (tokenString string, goError errorcode.Error) {
	expireAt := time.Now().Add(24 * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, authClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   user.Username,
			ExpiresAt: expireAt,
		},
		UserID: user.ID,
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		goError = helper.ErrorHandle(errorcode.ErrorService, errorcode.GenerateTokenError, err.Error())
		return "", goError
	}

	return
}
