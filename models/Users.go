package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(30)"`
	Email    string `gorm:"type:varchar(100)"`
	Password string `json:"-"`
}
