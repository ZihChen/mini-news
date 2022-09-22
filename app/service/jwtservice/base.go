package jwtservice

import (
	"github.com/golang-jwt/jwt"
	"mini-news/app/model"
	"sync"
)

type Interface interface {
	GenerateToken(user model.User) (token string, err error)
	ValidateToken(tokenString string) (int, string, error)
}

type service struct{}

var singleton *service
var once sync.Once

var jwtKey = []byte("Qt6/brgIDg00xD1cpBdfGJzMPS4x+nEZZ2lBSKqlqo7UOl8bWuMhEYnoVRND7HwQ\nm3jUqJXBooDKayuFZOhMiGj6Cj+K1HGE5NkPXo72koaIftkV+nYWmdlS0gRhJSg2\n7tQKRrI11OVuTWZW9AXfJs0PvHie8V183TCjNYIdOesyRSyxpKoItGaa5lKPXnS/\n840KoEnCFY3eqGSkavRvplljXH8WuWKV9CtPFw==")

type authClaims struct {
	jwt.StandardClaims
	UserID int `json:"userId"`
}

func NewService() Interface {
	once.Do(func() {
		singleton = &service{}
	})
	return singleton
}
