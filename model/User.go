package model

import (
	"gin_demo/utils/errmsg"
	"github.com/jinzhu/gorm"
)

//结构体值类型
type User struct {
	Model
	Username string `gorm:"type:varchar(100);comment:'昵称'" json:"username" label:"昵称"`
	Avatar   string `gorm:"type:varchar(250);comment:'头像'" json:"avatar" label:"头像"`
	Email    string `gorm:"type:varchar(100);not null;comment:'电子邮箱'" json:"email" validate:"required" label:"邮箱"`
	Password string `gorm:"type:varchar(20);not null;comment:'密码'" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Confirm  string `gorm:"type:varchar(20);not null;comment:'确认密码'" json:"confirm" validate:"required,min=6,max=20" label:"密码"`
	Phone    string `gorm:"type:varchar(20);comment:'手机号码'" json:"phone"`
	Role     int    `gorm:"type:int;default:1;comment:'权限 1:管理员 2:阅读者'" json:"role" label:"角色权限"`
}

type JsonUser struct {
	User
}

//查询用户是否存在 并返回
func CheckUser(email string, id int, username string) (data User, error interface{}) {
	var user User
	//First 查出第一个参数
	if email != "" {
		Db.Where("email = ?", email).First(&user)
	} else if id != 0 {
		Db.Where("id = ?", id).First(&user)
	} else if username != "" {
		Db.Where("username = ?", username).First(&user)
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
func GetUserList(pageSize int, PageNum int) ([]User, int) {
	var users []User
	var total int // 总数
	//一页多少个
	err := Db.Limit(pageSize).Offset((PageNum - 1) * pageSize).Find(&users).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return users, total
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
