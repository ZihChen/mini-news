package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"mini-news/app/handler/user_handler"
	"mini-news/app/middleware/jwt"
)

func LoadingRouter(r *gin.Engine) {

	userHandler := user_handler.NewsHandler()

	api := r.Group("/api")
	{
		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		user := api.Group("/user")
		{
			user.POST("/create", userHandler.CreateUser)
			user.POST("/login", userHandler.Login)
		}

		api.Use(jwt.VerifyToken)
		{

		}
	}
}
