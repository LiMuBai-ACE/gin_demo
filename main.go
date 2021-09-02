package main

import (
	"gin_demo/model"
	"gin_demo/routes"
)

func main() {
	//引用数据库 自动创建表
	model.InitDB()
	routes.InitRouter()
}
