package consts

const (
	UsernameRegexRule = `[a-zA-Z0-9_]{2,30}`
	PasswordRegexRule = `[a-zA-Z0-9]{6,20}`
)

const (
	RedisToken = "Token:"
)

const (
	TokenExpireTime = 3600 * 24 // 一天
)

const (
	BearerToken = "Bearer"
	BasicToken  = "Basic"
)
