package jwtservice

import (
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
	"mini-news/app/global/consts"
	"mini-news/app/global/errorcode"
	"mini-news/app/global/helper"
	"mini-news/app/global/settings"
	"mini-news/app/model"
	"strings"
	"time"
)

func (s *Service) GenerateToken(user model.User) (tokenMeta TokenMeta, goError errorcode.Error) {
	expireAt := time.Now().Add(24 * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, authClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   user.Username,
			ExpiresAt: expireAt,
		},
		UserID: user.ID,
	})

	tokenString, err := token.SignedString([]byte(settings.Config.JwtConfig.Key))
	if err != nil {
		goError = helper.ErrorHandle(errorcode.ErrorService, errorcode.GenerateTokenError, err.Error())
		return tokenMeta, goError
	}

	tokenMeta.Uuid = func() string {
		generator, _ := uuid.NewV4()
		return generator.String()
	}()
	tokenMeta.Signature = func() string {
		return strings.Split(tokenString, ".")[2]
	}()
	tokenMeta.TokenType = consts.BearerToken
	tokenMeta.ExpireAt = expireAt
	tokenMeta.AccessToken = tokenString

	return
}
