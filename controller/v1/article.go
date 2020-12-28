package v1

import (
	"gin_demo/model"
	"gin_demo/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//返回错误
func Msg(code int, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//添加文章
func AddArticle(c *gin.Context) {
	var article model.Article

	_ = c.ShouldBindJSON(&article)
	if article.Category.Name == "" {
		Msg(500, "分类不能为空!", c)
		return
	}
	if article.Title == "" {
		Msg(500, "标题不能为空!", c)
		return
	}
	if article.Content == "" {
		Msg(500, "文章内容不能为空!", c)
		return
	}
	if article.Desc == "" {
		Msg(500, "文章介绍不能为空!", c)
		return
	}
	if article.Img == "" {
		Msg(500, "文章封面不能为空!", c)
		return
	}

	//验证title是否重复
	data, _ := model.CheckArt(0, article.Title)
	if data.ID > 0 {
		Msg(500, "标题重复,请重新编辑标题!", c)
		return
	}

	code = model.CreateArt(&article)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrmsg(code),
	})
}

//查询单个文章
func GetArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	data, _ := model.CheckArt(id, "")
	if data.ID == 0 {
		Msg(500, "抱歉,未找到文章,该文章已关闭,或者被删除", c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})
}

//查询文章列表
func GetArticleList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	cid, _ := strconv.Atoi(c.Query("cid"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data := model.GetArtList(pageSize, pageNum, cid)
	code := errmsg.SUCCSE
	c.JSON(http.StatusOK, gin.H{
		"status":     code,
		"data":       data,
		"pageNum":    pageNum,
		"pageSize":   pageSize,
		"totalCount": len(model.GetArtList(-1, -1, cid)),
		"msg":        errmsg.GetErrmsg(code),
	})
}

//修改文章
func EditArt(c *gin.Context) {
	var article model.Article
	_ = c.ShouldBindJSON(&article)

	//验证标题是否重复
	data, _ := model.CheckArt(0, article.Title)
	if data.ID > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"msg":    "标题名称已存在,请更换标题名称",
		})
		return
	}

	code = model.EditArt(&article)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrmsg(code),
	})
}

//删除文章
func DeleteArt(c *gin.Context) {
	var article model.Article
	_ = c.ShouldBindJSON(&article)

	data, _ := model.CheckArt(int(article.ID), "")
	if data.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": errmsg.ERROR,
			"msg":  "文章不存在",
		})
		return
	}
	code = model.DeleteArt(int(article.ID))
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrmsg(code),
	})
}
