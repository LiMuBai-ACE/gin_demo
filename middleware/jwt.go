package middleware

import (
	"gin_demo/model"
	"gin_demo/utils"
	"gin_demo/utils/errmsg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var JwtKey = []byte(utils.JwtKey)

type MyClaims struct {
	jwt.StandardClaims
	Email string `json:"email"`
}

//生成token
func SetToken(email string) (string, int) {
	//过期时间 十小时
	expireTime := time.Now().Add(10 * time.Hour)
	SetClaims := MyClaims{
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), // 过期时间
			Issuer:    "ginblog",
		},
		email,
	}
	//签发的方法,生成的结构体 最后返回token
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return "", errmsg.ERROR
	}
	return token, errmsg.SUCCSE
}

//验证token
func CheckToken(token string) (*MyClaims, interface{}) {
	var claims MyClaims

	setToken, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (i interface{}, e error) {
		return JwtKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errmsg.ERROR_TOKEN_WRONG
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return nil, errmsg.ERROR
			} else {
				return nil, errmsg.ERROR
			}
		}
	}
	if setToken != nil {
		if key, ok := setToken.Claims.(*MyClaims); ok && setToken.Valid {
			return key, errmsg.SUCCSE
		} else {
			return nil, errmsg.ERROR_TOKEN_WRONG
		}
	}
	return nil, errmsg.ERROR_TOKEN_WRONG
}

var code int

//jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHerder := c.Request.Header.Get("Authorization") // 拿到写入的请求头token 进行验证
		if tokenHerder == "" {
			code = errmsg.ERROR_TOKEN_EXIST
			c.JSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  errmsg.GetErrmsg(code),
			})
			c.Abort()
			return
		}

		checkToken := strings.SplitN(tokenHerder, " ", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  errmsg.GetErrmsg(code),
			})
			c.Abort()
			return
		}
		//未能正常解析
		key, tCode := CheckToken(checkToken[1])
		// 不等于nil 返回不正确 则直接返回
		if tCode != nil && tCode != 200 {
			code = errmsg.ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  errmsg.GetErrmsg(code),
			})
			c.Abort()
			return
		}

		// token过期
		if time.Now().Unix() > key.ExpiresAt {
			code = errmsg.ERROR_TOKEN_RUNIME
			c.JSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  errmsg.GetErrmsg(code),
			})
			c.Abort()
			return
		}

		// 用户不存在
		data, _ := model.CheckUser(key.Email, 0, "")
		if data.ID == 0 {
			code = errmsg.ERROR_USER_NOT_EXIST
			c.JSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  errmsg.GetErrmsg(code),
			})
			c.Abort()
			return
		}
		c.Set("email", key.Email)
		c.Next()
	}
}
