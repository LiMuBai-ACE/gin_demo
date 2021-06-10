package routes

import (
	v1 "gin_demo/controller/v1"
	"gin_demo/middleware"
	"gin_demo/utils"
	"gin_demo/weather"
	"github.com/gin-gonic/gin"
)

//routerV1 接口的第一个版本
func InitRouter() {
	gin.SetMode(utils.Data.Server.Mode)
	r := gin.New()
	r.Use(middleware.Loggoer())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	Auth := r.Group("api/v1")
	Auth.Use(middleware.JwtToken()) // 需要验证的
	{
		//	用户模块的路由接口
		Auth.GET("user/info", v1.UserInfo)
		Auth.GET("users", v1.GetUserList)
		Auth.GET("user", v1.GetUser)
		Auth.POST("user/delete", v1.DeleteUser)
		Auth.POST("user/edit", v1.EditUser)
		//	分类模块的路由接口
		Auth.POST("category/delete", v1.DeleteCategory)
		Auth.POST("category", v1.Category)
		// 标签模块
		Auth.POST("tag/delete", v1.DeleteTag)
		Auth.POST("tag", v1.Tag)
		//	文章模块的路由接口
		Auth.POST("article", v1.Article)
		Auth.POST("article/delete", v1.DeleteArt)

		//	上传文件
		Auth.POST("upload", v1.UpLoad)
	}

	router := r.Group("api/v1")
	{
		//用户
		router.POST("user/add", v1.AddUser)
		router.POST("login", v1.Login)
		//分类
		router.GET("categorys", v1.GetCategoryList)
		//标签
		router.GET("tags", v1.GetCTagList)
		//文章
		router.GET("article", v1.GetArticle)
		router.GET("update/reading", v1.UpdateReading)
		router.GET("articles", v1.GetArticleList)
	}

	//关于天气的
	wrouter := r.Group("/")
	{
		wrouter.GET("/weather", weather.ShowWeather)
		wrouter.GET("/citylist", weather.ShowCityList)
		wrouter.GET("/weather/status", weather.ShowStatus)
	}
	panic(r.Run(utils.Data.Server.Port))
}
