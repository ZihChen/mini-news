package errorcode

const (
	ErrorHandler    ErrComponent = "handler"
	ErrorService    ErrComponent = "service"
	ErrorBusiness   ErrComponent = "business"
	ErrorRepository ErrComponent = "repository"
	ErrorDatabase   ErrComponent = "database"
	ErrorMiddleware ErrComponent = "middleware"
)

/* API Response error*/
const (
	/* http error*/
	BadRequest       ResponseErrType = "Bad Request"
	NotFoundResource ResponseErrType = "Not Found"
	AlreadyExists    ResponseErrType = "Already Exists"

	/* Request validate error */
	JsonBindError          ResponseErrType = "JSON Bind Error"
	StructToMapError       ResponseErrType = "Struct To Map Error"
	InvalidParamError      ResponseErrType = "Invalid Param Error"
	JsonMarshalError       ResponseErrType = "Json Marshal Error"
	RegexMatchError        ResponseErrType = "Regex Match Error"
	UsernameRegexRuleError ResponseErrType = "Username Regex Rule Error"
	PasswordRegexRuleError ResponseErrType = "Password Regex Rule Error"

	CheckUserExistError        ResponseErrType = "Check User Exist Error"
	UserCreateError            ResponseErrType = "User Create Error"
	UserPasswordIncorrectError ResponseErrType = "User Password Incorrect Error"
	GenerateTokenError         ResponseErrType = "Generate Token Error"
	GetHeaderTokenError        ResponseErrType = "Get Header Token Error"
	ParseWithClaimsError       ResponseErrType = "Parse with Claims Error"
	InvalidTokenError          ResponseErrType = "Invalid Token Error"
	GetAbmediaNewsError        ResponseErrType = "Get Abmedia News Error"
	VisitAbmediaNewsError      ResponseErrType = "Visit Abmedia News Error"
	GetArticleByTypeNamesError ResponseErrType = "Get Article by Type Names Error"
	InsertArticleByMapsError   ResponseErrType = "Insert Article by Maps Error"
	GetUserError               ResponseErrType = "Get User Error"
	UserAlreadyExist           ResponseErrType = "User Already Exist"
	HashPasswordError          ResponseErrType = "Hash Password Error"

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
