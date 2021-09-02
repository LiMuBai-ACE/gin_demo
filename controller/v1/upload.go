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
	url, code := model.UpLoadLocalFile(file)
	m := make(map[string]interface{})
	m["url"] = "http://" + c.Request.Host + "/" + url
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrmsg(code),
		"data": m,
	})
}
