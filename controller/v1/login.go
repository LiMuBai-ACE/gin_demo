package v1

import (
	"fmt"
	"gin_demo/middleware"
	"gin_demo/model"
	"gin_demo/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// 登录
func Login(c *gin.Context) {
	var data model.User
	c.ShouldBindJSON(&data)

	var token string
	var code int
	fmt.Println(data.Password)
	//验证密码是否通过
	code = model.CheckLogin(data.Email, data.Password)
	if code == errmsg.SUCCSE {
		token, code = middleware.SetToken(data.Email)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   errmsg.GetErrmsg(code),
		"token": token,
	})
}

// 利用token查询用户信息
func UserInfo(c *gin.Context) {
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

	key, tCode := middleware.CheckToken(checkToken[1])
	if tCode != nil {
		code = errmsg.ERROR_TOKEN_WRONG
		c.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  errmsg.GetErrmsg(code),
		})
		c.Abort()
		return
	}
	data, _ := model.CheckUser(key.Email, 0, "")
	if data.ID > 0 {
		code = errmsg.SUCCSE
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"data": data,
			"msg":  errmsg.GetErrmsg(code),
		})

	} else {
		code = errmsg.ERROR_USER_NOT_EXIST
		c.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  errmsg.GetErrmsg(code),
		})
	}
	return

}
