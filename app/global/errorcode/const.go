package errorcode

const (
	ErrorHandler    ErrComponent = "handler"
	ErrorService    ErrComponent = "service"
	ErrorBusiness   ErrComponent = "business"
	ErrorRepository ErrComponent = "repository"
	ErrorDatabase   ErrComponent = "database"
	ErrorRedis      ErrComponent = "redis"
	ErrorMiddleware ErrComponent = "middleware"
	ErrorCronJob    ErrComponent = "cronjob"
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

	ArticleTagCreateError      ResponseErrType = "Article Tag Create Error"
	TagsCountNotMatchError     ResponseErrType = "Tags Count Not Match Error"
	TagCreateError             ResponseErrType = "Tag Create Error"
	GetTagsCountError          ResponseErrType = "Get Tags Count Error"
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
	FunctionNameNotFound       ResponseErrType = "Function Name Not Found"

	/* Database error*/
	DBConnectError     ResponseErrType = "DB Connect Error"
	DBConnectPoolError ResponseErrType = "DB ConnectPool Error"

	/* Redis error*/
	RedisHashSetError      ResponseErrType = "Redis Hash Set Error"
	RedisSetError          ResponseErrType = "Redis Set Error"
	CheckHashKeyExistError ResponseErrType = "Check Hash Key Exist Error"
	RedisSetExpireError    ResponseErrType = "Redis Set Expire Error"
)

const (
	CheckDBConnectError   string = "Check DB Connect Error: %v"
	CheckConnectPoolError string = "Check ConnectPool Error: %v"
	PingDBError           string = "Ping DB Error: %v"
	RedisConnectError     string = "Redis Connect Error: %v"
	PingRedisError        string = "Ping Redis Error: %v"
	ReadFileError         string = "Read File Error: %v"
	YamlUnmarshalError    string = "Yaml Unmarshal Error: %v"

	UserTableMigrateError          string = "User Table Migrate Error: %v"
	CryptoArticleTableMigrateError string = "CryptoArticle Table Migrate Error: %v"
	TagTableMigrateError           string = "Tag Table Migrate Error: %v"
	ArticleTagTableMigrateError    string = "ArticleTag Table Migrate Error: %v"
)
