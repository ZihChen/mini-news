package errorcode

type ErrComponent string

const (
	ErrorHandler    ErrComponent = "handler"
	ErrorService    ErrComponent = "service"
	ErrorBusiness   ErrComponent = "business"
	ErrorRepository ErrComponent = "repository"
)

type ResponseErrType string

const (
	/* http error*/
	BadRequest       ResponseErrType = "Bad Request"
	NotFoundResource ResponseErrType = "Not Found"
	AlreadyExists    ResponseErrType = "Already Exists"

	JsonBindError     ResponseErrType = "JSON Bind Error"
	StructToMapError  ResponseErrType = "Struct To Map Error"
	InvalidParamError ResponseErrType = "Invalid Param Error"

	UserCreateError ResponseErrType = "User Create Error"
)

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
