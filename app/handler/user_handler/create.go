package user_handler

import (
	"github.com/gin-gonic/gin"
	"mini-news/app/global/response"
	"mini-news/app/global/structs/request"
	"mini-news/app/global/validator"
	"net/http"
)

func (h *Handler) CreateUser(c *gin.Context) {
	option := request.CreateUserOption{}
	if err := c.ShouldBindJSON(&option); err != nil {
		response.WrapContext(c).Error(1001, err.Error())
		return
	}

	if err := validator.ValidateParams(option); err != nil {
		response.WrapContext(c).Error(1002, err.Error())
		return
	}

	if err := h.userBusiness.CreateUser(option); err != nil {
		response.WrapContext(c).Error(1003, err.Error())
	}

	response.WrapContext(c).Success(http.StatusOK, "")
}
