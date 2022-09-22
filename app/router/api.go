package router

import (
	"github.com/gin-gonic/gin"
	"mini-news/app/handler/user_handler"
)

func LoadingRouter(r *gin.Engine) {

	userHandler := user_handler.NewsHandler()

	api := r.Group("/api")
	{
		user := api.Group("/user")
		{
			user.POST("/create", userHandler.CreateUser)
		}
	}
}
