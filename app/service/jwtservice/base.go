package jwtservice

import (
	"github.com/golang-jwt/jwt"
	"mini-news/app/global/errorcode"
	"mini-news/app/model"
	"sync"
)

type Interface interface {
	GenerateToken(user model.User) (token string, goError errorcode.Error)
	ValidateToken(tokenString string) (id int, username string, goError errorcode.Error)
}

type Service struct{}

var singleton *Service
var once sync.Once

type authClaims struct {
	jwt.StandardClaims
	UserID int `json:"userId"`
}

func NewService() Interface {
	once.Do(func() {
		singleton = &Service{}
	})
	return singleton
}
