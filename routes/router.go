package routes

import (
	v1 "gin_demo/controller/v1"
	"gin_demo/middleware"
	"gin_demo/utils"
	"github.com/gin-gonic/gin"
)

//routerV1 接口的第一个版本
func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	Auth := r.Group("api/v1")
	Auth.Use(middleware.JwtToken()) // 需要验证的
	{
		//	用户模块的路由接口
		Auth.POST("user/add", v1.AddUser)
		Auth.GET("users", v1.GetUserList)
		Auth.GET("user", v1.GetUser)
		Auth.POST("user/delete", v1.DeleteUser)
		Auth.POST("user/edit", v1.EditUser)
		//	分类模块的路由接口
		Auth.POST("category/add", v1.AddCategory)
		//Auth.GET("category", v1.GetCategory) // 无需单个查询 直接在列表修改

		Auth.POST("category/delete", v1.DeleteCategory)
		Auth.POST("category/edit", v1.EditCategory)
		//	文章模块的路由接口
		Auth.POST("article/add", v1.AddArticle)

		Auth.POST("article/edit", v1.EditArt)
		Auth.POST("article/delete", v1.DeleteArt)
	}

	router := r.Group("api/v1")
	{
		router.GET("categorys", v1.GetCategoryList)
		router.GET("article", v1.GetArticle)
		router.GET("articles", v1.GetArticleList)
		router.POST("login", v1.Login)
	}
	panic(r.Run(utils.HttpPort))
}
