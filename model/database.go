package model

import (
	"fmt"
	"gin_demo/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

// 创建db实例值
var Db *gorm.DB
var Err error

//同目录直接引用 该包里面的命名就好了

//连接数据库 创建db实例
func InitDB() {
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)

	Db, Err = gorm.Open(utils.Db, args)
	if Err != nil {
		panic("failed to connect database, err" + Err.Error())
	}

	// 禁用默认表名的复数形式,如果置为true,则`User`的默认表名是`user`
	Db.SingularTable(true)

	//创建数据表
	Db.AutoMigrate(&User{}, &Article{}, &Category{})

	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	Db.DB().SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	Db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。 10秒
	Db.DB().SetConnMaxLifetime(10 * time.Second)

	//defer db.Close()
}
