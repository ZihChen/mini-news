package jwtservice

import (
	"github.com/golang-jwt/jwt"
	"mini-news/app/model"
	"time"
)

func (s *service) GenerateToken(user model.User) (tokenString string, err error) {
	expireAt := time.Now().Add(24 * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, authClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   user.Username,
			ExpiresAt: expireAt,
		},
		UserID: user.ID,
	})

	tokenString, err = token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return
}