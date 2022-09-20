package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func LoadingRouter(r *gin.Engine) {

	api := r.Group("/api")

	api.GET("/", func(context *gin.Context) {
		fmt.Println("test")
	})
}
