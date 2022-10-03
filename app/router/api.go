package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"mini-news/app/business/cryptoarticle_business"
	"mini-news/app/handler/user_handler"
	"mini-news/app/middleware/jwt"
	"mini-news/app/middleware/logwriter"
)

func LoadingRouter(r *gin.Engine) {

	userHandler := user_handler.NewsHandler()

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
			api.GET("/test", func(c *gin.Context) {
				b := cryptoarticle_business.NewBusiness()
				b.SyncAbmediaArticles()
			})
		}
	}
}
