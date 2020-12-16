package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username      string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:type:varchar(20);not null" json:"password"`
	Phone string `gorm:"type:type:varchar(20);not null" json:"phone"`
	Role  string `gorm:"type:type:int" json:"role"`
}
