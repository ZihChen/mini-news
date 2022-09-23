package response

import (
	"github.com/gin-gonic/gin"
	"mini-news/app/global/errorcode"
)

type Wrapper struct {
	ctx *gin.Context
}

func WrapContext(ctx *gin.Context) *Wrapper {
	return &Wrapper{
		ctx: ctx,
	}
}

func (w *Wrapper) Json(response *Response) {
	w.ctx.JSON(200, response)
}

func (w *Wrapper) Success(statusCode int, data interface{}) {
	response := New()
	response.StatusCode = statusCode
	response.Data = data
	w.Json(response)
}

func (w *Wrapper) Error(statusCode int, message errorcode.Error) {
	response := New()
	response.StatusCode = statusCode
	response.Message = message
	w.Json(response)
}
