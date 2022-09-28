package jwt

import (
	"github.com/gin-gonic/gin"
	"mini-news/app/global/errorcode"
	"mini-news/app/global/helper"
	"mini-news/app/global/response"
	"mini-news/app/service/jwtservice"
	"net/http"
	"strings"
)

func VerifyToken(c *gin.Context) {
	token, ok := getToken(c)
	if !ok {
		goErr := helper.ErrorHandle(errorcode.ErrorMiddleware, errorcode.GetHeaderTokenError, "")
		response.WrapContext(c).Error(http.StatusOK, goErr)
		c.Abort()
		return
	}

	jwtService := jwtservice.NewService()

	id, username, goErr := jwtService.ValidateToken(token)
	if goErr != nil {
		response.WrapContext(c).Error(http.StatusOK, goErr)
		c.Abort()
		return
	}

	c.Set("id", id)
	c.Set("username", username)

	c.Next()

}

func getToken(c *gin.Context) (string, bool) {
	authValue := c.GetHeader("Authorization")

	arr := strings.Split(authValue, " ")
	if len(arr) != 2 {
		return "", false
	}

	authType := strings.Trim(arr[0], "\n\r\t")
	if strings.ToLower(authType) != strings.ToLower("Bearer") {
		return "", false
	}

	return strings.Trim(arr[1], "\n\t\r"), true
}
