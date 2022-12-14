package tag_handler

import (
	"github.com/gin-gonic/gin"
	"mini-news/app/global/errorcode"
	"mini-news/app/global/helper"
	"mini-news/app/global/response"
	"mini-news/app/global/structs/request"
	"mini-news/app/global/validator"
	"net/http"
)

func (h *Handler) CreateTag(c *gin.Context) {
	userID, _ := c.Get("id")
	option := request.CreateTagOption{
		UserID: userID.(int),
	}
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

	if goErr := h.tagBusiness.CreateTag(option); goErr != nil {
		response.WrapContext(c).Error(http.StatusOK, goErr)
		return
	}

	response.WrapContext(c).Success(http.StatusOK, "")
}
