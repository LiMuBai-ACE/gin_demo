package v1

import (
	"gin_demo/model"
	"gin_demo/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//添加分类
func AddCategory(c *gin.Context) {
	var category model.Category

	c.ShouldBindJSON(&category)
	//	验证不能为空
	if category.Name == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"msg":    "分类名称不能为空",
		})
		return
	}
	//验证name是否重复
	data, _ := model.CheckCategory(0, category.Name)
	if data.ID > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"msg":    "分类名称已存在,请更换分类名称",
		})
		return
	}

	code = model.CreateCategory(&category)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrmsg(code),
	})
}

//查询单个分类
func GetCategory(c *gin.Context) {

}

//查询分类
func GetCategoryList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, total := model.GetCategoryList(pageSize, pageNum)
	code := errmsg.SUCCSE
	c.JSON(http.StatusOK, gin.H{
		"status":     code,
		"data":       data,
		"pageNum":    pageNum,
		"pageSize":   pageSize,
		"totalCount": total,
		"msg":        errmsg.GetErrmsg(code),
	})
}

//修改分类
func EditCategory(c *gin.Context) {
	var category model.Category
	c.ShouldBindJSON(&category)

	//验证name是否重复
	data, _ := model.CheckCategory(0, category.Name)
	if data.ID > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"msg":    "分类名称已存在,请更换分类名称",
		})
		return
	}

	code = model.EditCategory(&category)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrmsg(code),
	})
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
