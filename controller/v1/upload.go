package v1

import (
	"gin_demo/model"
	"github.com/gin-gonic/gin"
)

func UpLoad(c *gin.Context) {
	var file model.File
	c.ShouldBindJSON(&file)
	model.UpLoadFile(file)
	//c.JSON(http.StatusOK, gin.H{
	//	"code": code,
	//	"msg":  errmsg.GetErrmsg(code),
	//	"url":  url,
	//})

}
