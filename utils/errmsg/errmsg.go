package errmsg

//错误信息中间件
const (
	SUCCSE = 200
	ERROR  = 500

	// code = 1000...用户模块的错误
	ERROR_USERNAME_USED  = 1001
	ERROR_PASSWORD_WRONG = 1002
	ERROR_USER_NOT_EXIST = 1003
	ERROR_TOKEN_EXIST    = 1004
	ERROR_TOKEN_RUNIME   = 1005
	ERROR_TOKEN_WRONG    = 1006
	ERROR_TOKEN_TYPE     = 1007
	ERROR_EMAIL_WRONG    = 1008

	// code = 2000...文章模块的错误
	// code = 3000...分类模块的错误

)

var codemsg = map[int]string{
	SUCCSE:               "ok",
	ERROR:                "error",
	ERROR_USERNAME_USED:  "用户已存在",
	ERROR_PASSWORD_WRONG: "密码错误",
	ERROR_USER_NOT_EXIST: "用户不存在",
	ERROR_TOKEN_EXIST:    "TOKEN不存在",
	ERROR_TOKEN_RUNIME:   "TOKEN已过期",
	ERROR_TOKEN_WRONG:    "TOKEN不正确",
	ERROR_TOKEN_TYPE:     "TOKEN格式错误",
	ERROR_EMAIL_WRONG:    "邮箱格式不正确,请传入正确的邮箱账号",
}

func GetErrmsg(code int) string {

	return codemsg[code]
}
