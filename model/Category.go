package model

import (
	"gin_demo/utils/errmsg"
	"github.com/jinzhu/gorm"
	"time"
)

type Category struct {
	ID        uint `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Name      string     `gorm:"type:varchar(20);not null" json:"name"`
}

//查询用户是否存在 并返回
func CheckCategory(id int, name string) (data Category, error interface{}) {
	var category Category
	//First 查出第一个参数
	if id != 0 {
		Db.Where("id = ?", id).First(&category)
	} else if name != "" {
		Db.Select("id").Where("username = ?", name).First(&category)
	} else {
		return category, "error"
	}
	return category, nil
}

//新增分类
func CreateCategory(data *Category) int {
	err := Db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//查询标签列表
func GetCategoryList(pageSize int, PageNum int) []Category {
	var category []Category
	//一页多少个
	err := Db.Limit(pageSize).Offset((PageNum - 1) * pageSize).Find(&category).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return category
}

//删除用户
func DeleteCategory(id int) int {
	var category Category
	//软删除
	err := Db.Where("id = ?", id).Delete(&category).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//修改用户
func EditCategory(id int, data *User) int {
	var category Category
	var categoryMaps = make(map[string]interface{})
	categoryMaps["username"] = data.Username
	categoryMaps["phone"] = data.Phone
	categoryMaps["role"] = data.Role
	err := Db.Model(&category).Where("id = ?", id).Updates(categoryMaps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
