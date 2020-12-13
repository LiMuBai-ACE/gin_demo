package v1

import (
	"fmt"
	"gin_demo/config"
	"gin_demo/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"math/rand"
	"net/http"
	"time"
)

func Reginster(ctx *gin.Context) {
	//使用DB实例
	DB := config.GetDB()
	//	获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	//	数据验证
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"cdoe": 422, "masg": "手机号必须为11位"})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"cdoe": 422, "masg": "手机号必须为6位"})
	}
	//如果名称没有传,给一个随机的字符串
	if len(name) == 0 {
		name = RandomString(10)
	}
	//打印传入的信息
	fmt.Println(name, telephone, password)

	//	判断手机号码是否存在
	if isTelephoneExist(DB, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"cdoe": 422, "masg": "用户已存在"})
		return
	}
	//	创建用户

	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	DB.Create(&newUser)

	//	返回结果
	ctx.JSON(200, gin.H{"mag": "注册成功"})

}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

//生成随机字符串
func RandomString(num int) string {
	var letters = []byte("asdfghjklzxcvbnmqwertyuiopASDFGHJKLZXCVBNMQWERTYUIOP")
	result := make([]byte, num)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
