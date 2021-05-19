package v1

import (
	"gin_demo/model"
	"gin_demo/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

//查询分类
func GetCategoryList(c *gin.Context) {
	data, total := model.GetCategoryList()

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
func Category(c *gin.Context) {
	var category model.Category
	c.ShouldBindJSON(&category)

	//	验证不能为空
	if category.Name == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "分类名称不能为空",
		})
		return
	}

	//验证name是否重复
	data, _ := model.CheckCategory(0, category.Name)
	if data.ID > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "分类名称已存在,请更换分类名称",
		})
		return
	}

	// 创建和修改
	if category.ID == 0 {
		code = model.CreateCategory(&category)
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errmsg.GetErrmsg(code),
		})
	} else {
		code = model.EditCategory(&category)
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errmsg.GetErrmsg(code),
		})
	}

}

//删除分类
func DeleteCategory(c *gin.Context) {
	var category model.Category
	c.ShouldBindJSON(&category)

	data, _ := model.CheckCategory(int(category.ID), "")
	if data.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": errmsg.ERROR,
			"msg":  "分类不存在",
		})
		return
	}
	code = model.DeleteCategory(int(category.ID))
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrmsg(code),
	})
}
