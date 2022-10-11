package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"mini-news/app/handler/tag_handler"
	"mini-news/app/handler/user_handler"
	"mini-news/app/middleware/jwt"
	"mini-news/app/middleware/logwriter"
)

func LoadingRouter(r *gin.Engine) {

	userHandler := user_handler.NewsHandler()
	tagHandler := tag_handler.NewHandler()

	api := r.Group("/api", logwriter.RequestLog)
	{
		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		user := api.Group("/user")
		{
			user.POST("/register", userHandler.RegisterUser)
			user.POST("/login", userHandler.Login)
		}

		api.Use(jwt.VerifyToken)
		{
			tag := api.Group("/tag")
			{
				tag.POST("/create", tagHandler.CreateTag)
			}
		}
	}
}
