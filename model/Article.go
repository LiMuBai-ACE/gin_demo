package model

import (
	"gin_demo/utils/errmsg"
	"github.com/jinzhu/gorm"
	"time"
)

//foreignkey 关联关系
type Article struct {
	Category  Category   `gorm:"foreignkey:Cid;"json:"category"`
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	Title     string     `gorm:"type:varchar(100);not null;comment:'文章标题'" json:"title"`
	Cid       uint       `gorm:"type:int;not null;comment:'分类id'" json:"cid"`
	Uid       uint       `gorm:"type:int;comment:'用户id'" json:"uid"`
	Desc      string     `gorm:"type:varchar(200);not null;comment:'文章简介'" json:"desc"`
	Content   string     `gorm:"type:longtext;not null;comment:'文章内容'" json:"content"`
	//Img       string     `gorm:"type:varchar(250);not null;comment:'文章图片'" json:"img"`
}

//查询文章是否存在 并返回
func CheckArt(id int, title string) (data Article, error interface{}) {
	var article Article
	//First 查出第一个参数
	if id != 0 {
		Db.Preload("Category").Where("id = ?", id).First(&article)
	} else if title != "" {
		Db.Preload("Category").Where("title = ?", title).First(&article)
	} else {
		return article, "error"
	}
	return article, nil
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
func GetArtList(pageSize int, PageNum int, cid int) ([]Article, int) {
	var article []Article
	var total int // 总数
	//一页多少个
	if cid != 0 {
		err := Db.Preload("Category").Where("cid = ?", cid).Limit(pageSize).Offset((PageNum - 1) * pageSize).Find(&article).Count(&total).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, 0
		}
	} else {
		// 预加载 Preload
		err := Db.Preload("Category").Limit(pageSize).Offset((PageNum - 1) * pageSize).Find(&article).Count(&total).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, 0
		}
	}
	return article, total
}

//修改文章
func EditArt(data *Article) int {
	var article Article
	var articleMaps = make(map[string]interface{})
	// 修改文章标签
	if data.Category.Name != "" && data.Cid != 0 {
		articleMaps["name"] = data.Category.Name
		articleMaps["cid"] = data.Cid
	}
	//修改内容
	if data.Title != "" {
		articleMaps["title"] = data.Title
	}
	if data.Desc != "" {
		articleMaps["desc"] = data.Desc
	}
	if data.Content != "" {
		articleMaps["content"] = data.Content
	}
	if data.Uid != 0 {
		articleMaps["uid"] = data.Uid
	}
	//if data.Img != "" {
	//	articleMaps["img"] = data.Img
	//}
	err := Db.Model(&article).Where("id = ?", data.ID).Updates(articleMaps).Error
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
