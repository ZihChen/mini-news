package response

import "github.com/gin-gonic/gin"

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

func (w *Wrapper) Success(data interface{}, message string) {
	response := New()
	response.Data = data
	response.Message = message
	w.Json(response)
}

func (w *Wrapper) Error(statusCode int, message string) {
	response := New()
	response.StatusCode = statusCode
	response.Message = message
	w.Json(response)
}
