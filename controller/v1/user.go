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
	var user model.User
	_ = c.ShouldBindJSON(&user)

	//验证邮箱账号
	if !utils.VerifyEmailFormat(user.Email) {
		code = errmsg.ERROR_EMAIL_WRONG
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"data":   user,
			"msg":    errmsg.GetErrmsg(code),
		})
		return
	}
	//根据邮箱查询用户是否存在
	data, _ := model.CheckUser(user.Email, 0, "")

	//用户已存在
	if data.ID > 0 {
		code = errmsg.ERROR_USERNAME_USED
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    errmsg.GetErrmsg(code),
		})
		return
	}
	user.Username = user.Email                     // 邮箱默认为用户昵称
	user.Password = utils.ScryptStr(user.Password) // 密码加密

	//添加数据
	code = model.CreateUser(&user)
	//数据创建成功
	if code == errmsg.SUCCSE {
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"data":   user,
			"msg":    errmsg.GetErrmsg(code),
		})
		return
	}
	//数据未创建成功
	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"data":   user,
			"msg":    errmsg.GetErrmsg(code),
		})
		return
	}
}

//查询单个用户
func GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	email := c.Query("email")

	if id != 0 {
		data, _ := model.CheckUser("", id, "")
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
				"code": code,
				"msg":  errmsg.GetErrmsg(code),
			})
		}
		return
	}
	if email != "" {
		data, _ := model.CheckUser(email, 0, "")
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
				"code": code,
				"msg":  errmsg.GetErrmsg(code),
			})
		}
		return
	}

	code = errmsg.ERROR_EMAIL_ID_email
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrmsg(code),
	})
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
	var params struct {
		model.User
		ID string
	}
	c.ShouldBindJSON(&params)
	// 验证用户是否存在
	user, _ := model.CheckUser("", 0, params.Username)
	if user.ID > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":   errmsg.GetErrmsg(errmsg.ERROR),
			"msg":    "该昵称已被使用,请更换其他昵称修改!",
			"status": errmsg.ERROR,
		})
		return
	}
	id, _ := strconv.Atoi(params.ID)
	fmt.Printf("%#v", id)
	code = model.EditUser(id, &params.User)
	if code == errmsg.ERROR {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errmsg.GetErrmsg(code),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errmsg.GetErrmsg(code),
		})
	}
}

//删除用户信息
func DeleteUser(c *gin.Context) {
	var idObj struct {
		ID string `json:"id"`
	}
	c.ShouldBindJSON(&idObj)
	id, _ := strconv.Atoi(idObj.ID)
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
	data, _ := model.CheckUser("", id, "")
	// 未查询用户
	if data.ID == 0 {
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
