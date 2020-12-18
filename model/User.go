package model

import (
	"gin_demo/utils/errmsg"
	"github.com/jinzhu/gorm"
)

//结构体值类型
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);" json:"username"`
	Email    string `gorm:"type:varchar(20);not null" json:"email"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Phone    string `gorm:"type:varchar(20);" json:"phone"`
	Role     int    `gorm:"type:int" json:"role"`
}

//根据id查询用户是否存在
func CheckUserId(id int) User {
	var user User
	//First 查出第一个参数
	Db.Where("id = ?", id).First(&user)
	return user
}

//查询用户是否存在
func CheckUser(email string) (code int) {
	var users User
	//First 查出第一个参数
	Db.Select("id").Where("email = ?", email).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCSE
}

//新增用户
func CreateUser(data *User) int {
	err := Db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//查询用户列表
func GetUserList(pageSize int, PageNum int) []User {
	var users []User
	//一页多少个
	err := Db.Limit(pageSize).Offset((PageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

//删除用户
func DeleteUser(id int) int {
	var user User
	//软删除
	err := Db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	return errmsg.SUCCSE
}
