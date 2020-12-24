package model

import (
	"gin_demo/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Article struct {
	Category
	Title   string `gorm:"type:varchar(100);not null;comment:'文章标题'" json:"title"`
	Cid     int    `gorm:"type:int;not null;comment:'分类id'" json:"cid"`
	Uid     int    `gorm:"type:int;not null;comment:'用户id'" json:"uid"`
	Desc    string `gorm:"type:varchar(200);comment:'文章简介'" json:"desc"`
	Content string `gorm:"type:longtext;comment:'文章内容'" json:"content"`
	Img     string `gorm:"type:varchar(100);comment:'文章图片'" json:"img"`
}

//查询文章是否存在 并返回
func CheckArt(id int, title string) (data Article, error interface{}) {
	var article Article
	//First 查出第一个参数
	if id != 0 {
		Db.Where("id = ?", id).First(&article)
	} else if title != "" {
		Db.Where("title = ?", title).First(&article)
	} else {
		return article, "error"
	}
	return article, nil
}

//查询分类下的文章
func GetCateArt(CId int) []Article {
	var article []Article
	err := Db.Where("cid", CId).Find(&article).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return article
}

//新增文章
func CreateArt(data *Article) int {
	err := Db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//查询文章列表
func GetArtList(pageSize int, PageNum int, cid int) []Article {
	var article []Article
	//一页多少个
	if cid != 0 {
		err := Db.Limit(pageSize).Offset((PageNum-1)*pageSize).Where("cid = ?", cid).Find(&article).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil
		}
	} else {
		err := Db.Limit(pageSize).Offset((PageNum - 1) * pageSize).Find(&article).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil
		}
	}
	return article
}

//修改文章
func EditArt(data *Article) int {
	var article Article
	var articleMaps = make(map[string]interface{})
	// 修改文章标签
	articleMaps["name"] = data.Name
	articleMaps["cid"] = data.Cid
	//修改内容
	articleMaps["title"] = data.Title
	articleMaps["desc"] = data.Desc
	articleMaps["content"] = data.Content
	articleMaps["img"] = data.Img
	err := Db.Model(&article).Where("id = ?", data.ID).Updates(article).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//删除文章
func DeleteArt(id int) int {
	var article Article
	//软删除
	err := Db.Where("id = ?", id).Delete(&article).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
