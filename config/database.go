package config

import (
	"fmt"
	"gin_demo/model"
	"github.com/jinzhu/gorm"
)

// 创建db实例值
var DB *gorm.DB

//连接数据库 创建db实例
func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "gin_demo"
	username := "root"
	password := "Li-123456"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset,
	)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database, err" + err.Error())
	}
	//创建数据表
	db.AutoMigrate(&model.User{})
	// 给DB赋值
	DB = db
	return db
}

//获取db实列  返回DB实例
func GetDB() *gorm.DB {
	//return DB
	return InitDB() // 另外一种写法
}
