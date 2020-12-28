package model

import (
	"gin_demo/utils"
	"gin_demo/utils/errmsg"
)

//登录验证
func CheckLogin(email string, password string) int {
	var user User
	Db.Where("email = ?", email).First(&user)

	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if utils.ScryptStr(password) != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 0 {
		return errmsg.ERROR_User_NO_RIGHT
	}
	return errmsg.SUCCSE
}
