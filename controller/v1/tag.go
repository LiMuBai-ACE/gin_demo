package v1

import (
	"gin_demo/model"
	"gin_demo/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

//查询分类
func GetCTagList(c *gin.Context) {
	data, total := model.GetTagList()

	m := make(map[string]interface{})
	m["list"] = data
	//m["pageNum"] = pageNum
	//m["pageSize"] = pageSize
	m["total"] = total

	code := errmsg.SUCCSE
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": m,
		"msg":  errmsg.GetErrmsg(code),
	})
}

// 添加-修改分类
func Tag(c *gin.Context) {
	var tag model.Tag
	c.ShouldBindJSON(&tag)

	//	验证不能为空
	if tag.Name == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "分类名称不能为空",
		})
		return
	}

	//验证name是否重复
	data, _ := model.CheckCategory(0, tag.Name)
	if data.ID > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "分类名称已存在,请更换分类名称",
		})
		return
	}

	// 创建和修改
	if tag.ID == 0 {
		code = model.CreateTag(&tag)
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errmsg.GetErrmsg(code),
		})
	} else {
		code = model.EditTag(&tag)
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errmsg.GetErrmsg(code),
		})
	}

}

//删除分类
func DeleteTag(c *gin.Context) {
	var tag model.Tag
	c.ShouldBindJSON(&tag)

	data, _ := model.CheckCategory(int(tag.ID), "")
	if data.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": errmsg.ERROR,
			"msg":  "分类不存在",
		})
		return
	}
	code = model.DeleteCategory(int(tag.ID))
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrmsg(code),
	})
}
