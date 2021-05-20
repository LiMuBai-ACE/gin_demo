package model

import (
	"gin_demo/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Tag struct {
	ID        uint    `gorm:"primary_key" json:"id"`
	CreatedAt MyTime  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt MyTime  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *MyTime `gorm:"column:deleted_at" sql:"index" json:"deleted_at"`
	Name      string  `gorm:"type:varchar(20);not null;comment:'标签名称'" json:"name"`
}

//查询分类是否存在 并返回
func CheckTag(id int, name string) (data Tag, error interface{}) {
	var tag Tag
	//First 查出第一个参数
	if id != 0 {
		Db.Where("id = ?", id).First(&tag)
	} else if name != "" {
		Db.Where("name = ?", name).First(&tag)
	} else {
		return tag, "error"
	}
	return tag, nil
}

//新增分类
func CreateTag(data *Tag) int {
	err := Db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//查询标签列表
func GetTagList() ([]Tag, int) {
	var tag []Tag
	var total int // 总数
	//一页多少个
	err := Db.Find(&tag).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return tag, total
}

//修改分类
func EditTag(data *Tag) int {
	var tag Tag
	var tagMaps = make(map[string]interface{})
	tagMaps["name"] = data.Name
	err := Db.Model(&tag).Where("id = ?", data.ID).Updates(tagMaps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//删除用户
func DeleteTag(id int) int {
	var tag Tag
	//软删除
	err := Db.Where("id = ?", id).Delete(&tag).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
