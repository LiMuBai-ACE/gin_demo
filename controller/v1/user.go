package v1

import (
	"gin_demo/model"
	"gin_demo/utils"
	"gin_demo/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var code int

//添加用户
func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)

	//验证邮箱账号
	if !utils.VerifyEmailFormat(data.Email) {
		code = errmsg.ERROR_EMAIL_WRONG
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"data":   data,
			"msg":    errmsg.GetErrmsg(code),
		})
		return
	}

	code = model.CheckUser(data.Email)

	data.Username = data.Email // 邮箱默认为用户昵称

	//数据创建成功
	if code == errmsg.SUCCSE {
		model.CreateUser(&data)
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"data":   data,
			"msg":    errmsg.GetErrmsg(code),
		})
	}

	//数据未创建成功
	if code == errmsg.ERROR {
		code = errmsg.ERROR_USERNAME_USED
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"data":   data,
			"msg":    errmsg.GetErrmsg(code),
		})
		return
	}
}

//查询单个用户
func GetUser(c *gin.Context) {

}

//查询用户列表
func GetUserList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data := model.GetUserList(pageSize, pageNum)
	code := errmsg.SUCCSE
	c.JSON(http.StatusOK, gin.H{
		"status":     code,
		"data":       data,
		"pageNum":    pageNum,
		"pageSize":   pageSize,
		"totalCount": len(model.GetUserList(-1, -1)),
		"msg":        errmsg.GetErrmsg(code),
	})
}

//修改用户信息
func EditUser(c *gin.Context) {
}

//删除用户信息
func DeleteUser(c *gin.Context) {

}
