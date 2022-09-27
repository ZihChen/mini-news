package errorcode

type ErrComponent string
type ResponseErrType string

type GoError struct {
	error
	ErrMsg       string
	Component    ErrComponent
	ResponseType ResponseErrType
}

type Error interface {
	GetErrorMsg() string
	SetErrMsg(errMsg string)
	SetErrComponent(Component ErrComponent)
	SetResponseErrType(ResponseType ResponseErrType)
}

func NewGoError() Error {
	return &GoError{}
}

func (e *GoError) GetErrorMsg() string {
	return e.ErrMsg
}

func (e *GoError) SetErrMsg(errMsg string) {
	e.ErrMsg = errMsg
}

func (e *GoError) SetErrComponent(component ErrComponent) {
	e.Component = component
}

func (e *GoError) SetResponseErrType(responseType ResponseErrType) {
	e.ResponseType = responseType
}
