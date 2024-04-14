package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `gorm:"primary_key"`
	Username string `gorm:"unique"`
	Password string
}
