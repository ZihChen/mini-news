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

// Login UserLogin
// @Summary 使用者登入
// @Tags User
// @Produce json
// @Router /api/user/login [POST]
func (h *Handler) Login(c *gin.Context) {
	option := request.UserLogin{}
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

	token, goErr := h.userBusiness.UserLogin(option)
	if goErr != nil {
		response.WrapContext(c).Error(http.StatusOK, goErr)
		return
	}

	response.WrapContext(c).Success(http.StatusOK, token)
}

// RegisterUser 註冊User
// @Summary 註冊User
// @Tags User
// @Produce json
// @Router /api/user/register [POST]
func (h *Handler) RegisterUser(c *gin.Context) {
	option := request.RegisterUserOption{}
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
