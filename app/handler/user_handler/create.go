package user_handler

import (
	"github.com/gin-gonic/gin"
	"mini-news/app/global/errorcode"
	"mini-news/app/global/helper"
	"mini-news/app/global/response"
	"mini-news/app/global/structs/request"
	"mini-news/app/global/validator"
	"net/http"
)

// CreateUser 創建User
// @Summary 創建User
// @Tags User
// @Produce json
// @Router /api/user/create [POST]
func (h *Handler) CreateUser(c *gin.Context) {
	option := request.CreateUserOption{}
	if err := c.ShouldBindJSON(&option); err != nil {
		goErr := helper.ErrorHandle(errorcode.ErrorHandler, errorcode.JsonBindError, err.Error())
		response.WrapContext(c).Error(http.StatusOK, goErr)
		return
	}

	if err := validator.ValidateParams(option); err != nil {
		goErr := helper.ErrorHandle(errorcode.ErrorHandler, errorcode.InvalidParamError, err.Error())
		response.WrapContext(c).Error(http.StatusOK, goErr)
		return
	}

	if goErr := h.userBusiness.CreateUser(option); goErr != nil {
		response.WrapContext(c).Error(http.StatusOK, goErr)
	}

	response.WrapContext(c).Success(http.StatusOK, "")
}
