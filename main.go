package main

import (
	"gin_demo/routes"
	"gin_demo/utils"
)

func main(){

	//引用数据库 自动创建表
	utils.InitDB()

	routes.InitRouter()
}