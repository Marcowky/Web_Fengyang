package model

import "github.com/jinzhu/gorm"

/* model/user.go */
type User struct {
	gorm.Model
	UserName    string `gorm:"varchar(20);not null;table:users_table"`
	PhoneNumber string `gorm:"varchar(20);not null;unique;table:users_table"`
	Password    string `gorm:"size:255;not null;table:users_table"`
	UserType    string `gorm:"varchar(20);not null;table:users_table"`
}

