package model

import (
	"gin_demo/utils/errmsg"
	"github.com/jinzhu/gorm"
)

//结构体值类型
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);comment:'昵称'" json:"username"`
	Email    string `gorm:"type:varchar(20);not null;comment:'电子邮箱'" json:"email"`
	Password string `gorm:"type:varchar(20);not null;comment:'密码'" json:"password"`
	Phone    string `gorm:"type:varchar(20);comment:'手机号码'" json:"phone"`
	Role     int    `gorm:"type:int;comment:'权限'" json:"role"`
}

//查询用户是否存在 并返回
func CheckUser(email string, id int, username string) (data User, error interface{}) {
	var user User
	//First 查出第一个参数
	if email != "" {
		Db.Select("id").Where("email = ?", email).First(&user)
	} else if id != 0 {
		Db.Where("id = ?", id).First(&user)
	} else if username != "" {
		Db.Select("id").Where("username = ?", username).First(&user)
	} else {
		return user, "error"
	}
	return user, nil
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

//修改用户
func EditUser(id int, data *User) int {
	var user User
	var userMaps = make(map[string]interface{})
	userMaps["username"] = data.Username
	userMaps["phone"] = data.Phone
	userMaps["role"] = data.Role
	err := Db.Model(&user).Where("id = ?", id).Updates(userMaps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
