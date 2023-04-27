package model

import "github.com/jinzhu/gorm"

/* model/user.go */
type User struct {
	gorm.Model
	UserName    string `gorm:"varchar(20);not null"`
	PhoneNumber string `gorm:"varchar(20);not null;unique"`
	Password    string `gorm:"size:255;not null"`
	// Avatar      string `gorm:"size:255;not null"`
	// Collects    Array  `gorm:"type:longtext"`
	// Following   Array  `gorm:"type:longtext"`
	// Fans        int    `gorm:"AUTO_INCREMENT"`
}

// 用于获取收藏的结构体，这里暂时不需要，故注释
// type UserInfo struct {
// 	ID       uint   `json:"id"`
// 	// Avatar   string `json:"avatar"`
// 	UserName string `json:"userName"`
// }
