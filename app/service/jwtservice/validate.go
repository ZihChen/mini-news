package jwtservice

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"mini-news/app/global/errorcode"
	"mini-news/app/global/helper"
	"mini-news/app/global/settings"
)

func (s *Service) ValidateToken(tokenString string) (id int, username string, goError errorcode.Error) {
	var claims authClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(settings.Config.JwtConfig.Key), nil
	})

	if err != nil {
		goError = helper.ErrorHandle(errorcode.ErrorService, errorcode.ParseWithClaimsError, err.Error())
		return 0, "", goError
	}

	if !token.Valid {
		goError = helper.ErrorHandle(errorcode.ErrorService, errorcode.InvalidTokenError, token.Raw)
		return 0, "", goError
	}

	id = claims.UserID
	username = claims.Subject

	return id, username, nil
}
