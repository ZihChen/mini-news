package response

import (
	"github.com/google/uuid"
	"mini-news/app/global/errorcode"
)

type Response struct {
	StatusCode int             `json:"status_code"`
	Message    errorcode.Error `json:"message"`
	Data       interface{}     `json:"data"`
	Meta       Meta            `json:"meta"`
	Errors     []ErrorItem     `json:"errors"`
}

type Meta struct {
	RequestId string `json:"request_id"`
}

type ErrorItem struct {
	Key   string `json:"key"`
	Value string `json:"error"`
}

func New() *Response {
	return &Response{
		StatusCode: 200,
		Message:    &errorcode.GoError{},
		Data:       nil,
		Meta: Meta{
			RequestId: func() string {
				uuid, _ := uuid.NewUUID()
				return uuid.String()
			}(),
		},
		Errors: []ErrorItem{},
	}
}
