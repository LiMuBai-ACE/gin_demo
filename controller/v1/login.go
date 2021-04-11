package v1

import (
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
	checkToken := strings.SplitN(tokenHerder, " ", 2)

	key, _ := middleware.CheckToken(checkToken[1])
	data, _ := model.CheckUser(key.Email, 0, "")
	c.JSON(http.StatusOK, gin.H{
		"code": errmsg.SUCCSE,
		"data": data,
		"msg":  errmsg.GetErrmsg(code),
	})

}
