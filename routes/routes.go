package routes

import (
	"gin_demo/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(){
	gin.SetMode(config.AppMode)
	r := gin.Default()
	router := r.Group("api/v1")
	{
		router.GET("/hello", func(c *gin.Context) {
				c.JSON(http.StatusOK,gin.H{
					"code":200,
					"mag": "成功",
				})
		})
	}
	panic(r.Run(config.HttpPort))
}