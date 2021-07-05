package v1

import (
	"gin_demo/model"
	"gin_demo/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpLoad(c *gin.Context) {
	var file model.File
	c.ShouldBindJSON(&file)
	url, code := model.UpLoadFile(file.File, file.FileName, file.Type)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrmsg(code),
		"url":  url,
	})

}
