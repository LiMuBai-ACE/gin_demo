package v1

import (
	"gin_demo/middleware"
	"gin_demo/model"
	"gin_demo/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
		"status": code,
		"msg":    errmsg.GetErrmsg(code),
		"token":  token,
	})
}
