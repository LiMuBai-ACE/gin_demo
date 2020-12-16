package utils

import (
	"fmt"
	"gin_demo/config"
	"gin_demo/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

// 创建db实例值
var db *gorm.DB
var err error

//同目录直接引用 该包里面的命名就好了

//连接数据库 创建db实例
func InitDB(){
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DbUser,
		config.DbPassWord,
		config.DbHost,
		config.DbPort,
		config.DbName,
	)

	db, err = gorm.Open(config.Db, args)
	if err != nil {
		panic("failed to connect database, err" + err.Error())
	}

	// 禁用默认表名的复数形式,如果置为true,则`User`的默认表名是`user`
	db.SingularTable(true)

	//创建数据表
	db.AutoMigrate(&model.User{},&model.Article{},&model.Category{})

	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	db.DB().SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。 10秒
	db.DB().SetConnMaxLifetime(10 * time.Second)

	//defer db.Close()
}

