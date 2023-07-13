package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"

	// "fmt"
	// "net/url"
	// "os"

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
func InitDB(mode string) *gorm.DB {
	var configPath string
	switch mode {
	case "online":
		configPath = "config/config.json"
	case "local":
		configPath = "config/localConfig.json"
	case "test":
		configPath = "config/testConfig.json"
	}

	var jsonData map[string]interface{}

	// 读取JSON文件内容
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic("failed to open database: " + err.Error())
	}

	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		panic("failed to open database: " + err.Error())
	}
	databaseConfig := jsonData["databaseConfig"].(map[string]interface{})
	// 设定驱动参数
	driverName := databaseConfig["driverName"].(string)
	user := databaseConfig["user"].(string)
	password := databaseConfig["password"].(string)
	host := databaseConfig["host"].(string)
	port := databaseConfig["port"].(string)
	database := databaseConfig["database"].(string)
	charset := databaseConfig["charset"].(string)
	loc := databaseConfig["loc"].(string)

	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
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
