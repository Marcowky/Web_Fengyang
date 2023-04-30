package common

import (
	"fmt"
	"net/url"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// InitDB() 数据库初始化
func InitDB() *gorm.DB {
	// 设定驱动参数
	driverName := "mysql"
	user := "root"
	password := "Wu123456" // 需修改
	host := "localhost"
	port := "3306"
	database := "web_fengyang" // 需修改为自己存储的数据库
	charset := "utf8"
	loc := "Asia/Shanghai"
	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		user,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc))

	// 连接数据库
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to open database: " + err.Error())
	}
	DB = db
	return db
}

// 数据库信息获取
func GetDB() *gorm.DB {
	return DB
}
