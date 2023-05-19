package common

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// 用于接收配置文件的结构体
type DatabaseConfig struct {
    Config struct {
        DriverName string `json:"driverName"`
        User       string `json:"user"`
        Password   string `json:"password"`
        Host       string `json:"host"`
        Port       string `json:"port"`
        Database   string `json:"database"`
        Charset    string `json:"charset"`
        Loc        string `json:"loc"`
    } `json:"databaseConfig"`
}

// InitDB() 数据库初始化
func InitDB() *gorm.DB {
	// 导入配置文件
	configFile, err := os.Open("config/config.json")
    if err != nil {
        panic("failed to open database: " + err.Error())
    }
    defer configFile.Close()
	// 解码JSON文件
    decoder := json.NewDecoder(configFile)
	var myConfig DatabaseConfig
    err = decoder.Decode(&myConfig)
    if err != nil {
        panic("failed to open database: " + err.Error())
    }

	// 设定驱动参数
	driverName := myConfig.Config.DriverName
	user := myConfig.Config.User
	password := myConfig.Config.Password
	host := myConfig.Config.Host
	port := myConfig.Config.Port
	database := myConfig.Config.Database
	charset := myConfig.Config.Charset
	loc := myConfig.Config.Loc
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
