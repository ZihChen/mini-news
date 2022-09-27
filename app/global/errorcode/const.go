package errorcode

const (
	ErrorHandler    ErrComponent = "handler"
	ErrorService    ErrComponent = "service"
	ErrorBusiness   ErrComponent = "business"
	ErrorRepository ErrComponent = "repository"
	ErrorDatabase   ErrComponent = "database"
)

/* API Response error*/
const (
	/* http error*/
	BadRequest       ResponseErrType = "Bad Request"
	NotFoundResource ResponseErrType = "Not Found"
	AlreadyExists    ResponseErrType = "Already Exists"

	/* Request validate error */
	JsonBindError     ResponseErrType = "JSON Bind Error"
	StructToMapError  ResponseErrType = "Struct To Map Error"
	InvalidParamError ResponseErrType = "Invalid Param Error"

	GetUserError               ResponseErrType = "Get User Error"
	UserCreateError            ResponseErrType = "User Create Error"
	UserPasswordIncorrectError ResponseErrType = "User Password Incorrect Error"
	GenerateTokenError         ResponseErrType = "Generate Token Error"

	/* Database error*/
	DBConnectError     ResponseErrType = "DB Connect Error"
	DBConnectPoolError ResponseErrType = "DB ConnectPool Error"
)

const (
	CheckDBConnectError   string = "Check DB Connect Error: %v"
	CheckConnectPoolError string = "Check ConnectPool Error: %v"
	PingDBError           string = "Ping DB Error: %v"

	UserTableMigrateError          string = "User Table Migrate Error: %v"
	CryptoArticleTableMigrateError string = "CryptoArticle Table Migrate Error: %v"
)
