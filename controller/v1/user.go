package v1

import (
	"fmt"
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

	data.Username = data.Email                     // 邮箱默认为用户昵称
	data.Password = utils.ScryptStr(data.Password) // 密码加密

	//数据创建成功
	if code == errmsg.SUCCSE {
		model.CreateUser(&data)
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"data":   data,
			"msg":    errmsg.GetErrmsg(code),
		})
		return
	}
	//数据未创建成功
	if code == errmsg.ERROR_USERNAME_USED {
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
	var idObj struct {
		Id string `json:"id"`
	}
	c.ShouldBindJSON(&idObj)
	id, _ := strconv.Atoi(idObj.Id)
	fmt.Println(id)
	//是否传入正确id
	if id == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":   errmsg.GetErrmsg(errmsg.ERROR),
			"msg":    "请传入正确的id",
			"status": errmsg.ERROR,
		})
		return
	}
	//查询id 查看用户是否存在
	user := model.CheckUserId(id)
	if user.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"msg":    errmsg.GetErrmsg(errmsg.ERROR_USER_NOT_EXIST),
			"status": errmsg.ERROR_USER_NOT_EXIST,
		})
		return
	}
	//删除是否成功
	code = model.DeleteUser(id)
	if code == errmsg.ERROR_USER_NOT_EXIST {
		c.JSON(http.StatusOK, gin.H{
			"msg":    errmsg.GetErrmsg(code),
			"status": code,
		})
		return
	}
	if code == errmsg.SUCCSE {
		c.JSON(http.StatusOK, gin.H{
			"msg":    errmsg.GetErrmsg(code),
			"status": code,
		})
		return
	}
}
