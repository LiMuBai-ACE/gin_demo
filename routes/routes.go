package routes

import (
	"gin_demo/controller/v1"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", v1.Reginster)
	return r
}
