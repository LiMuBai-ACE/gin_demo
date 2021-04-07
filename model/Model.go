package model

import "time"

type Model struct {
	ID        uint       `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}
