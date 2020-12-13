package main

import (
	"gin_demo/common"
	_ "github.com/go-sql-driver/mysql"

	//引入结构体
	"github.com/gin-gonic/gin"
)

func main() {
	db := common.GetDB()
	defer db.Close() // 延迟关闭数据库
	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run()) // 监听并在 0.0.0.0:8080 上启动服务
}
