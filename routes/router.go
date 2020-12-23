package routes

import (
	v1 "gin_demo/controller/v1"
	"gin_demo/utils"
	"github.com/gin-gonic/gin"
)

//routerV1 接口的第一个版本
func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	routerV1 := r.Group("api/v1")
	{
		//	用户模块的路由接口
		routerV1.POST("user/add", v1.AddUser)
		routerV1.GET("users", v1.GetUserList)
		routerV1.GET("user", v1.GetUser)
		routerV1.POST("user/delete", v1.DeleteUser)
		routerV1.POST("user/edit", v1.EditUser)
		//	分类模块的路由接口
		routerV1.POST("category/add", v1.AddCategory)
		routerV1.GET("categorys", v1.GetCategoryList)
		//routerV1.GET("category", v1.GetCategory) // 无需单个查询 直接在列表修改

		routerV1.POST("category/delete", v1.DeleteCategory)
		routerV1.POST("category/edit", v1.EditCategory)
		//	文章模块的路由接口
	}
	panic(r.Run(utils.HttpPort))
}
