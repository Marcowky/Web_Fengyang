package main

import (
	"Web_Fengyang_Server/common"
	"Web_Fengyang_Server/model"
	"Web_Fengyang_Server/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	// "net/http"
)

func main() {
	// 获取初始化的数据库
	db := common.InitDB()
	// 迁移user数据表
	db.AutoMigrate(model.User{})
	// 延迟关闭数据库
	defer db.Close()
	// 创建路由引擎
	r := gin.Default()
	// 配置静态文件路径
	// r.StaticFS("/images", http.Dir("./static/images"))
	// 启动路由
	routes.CollectRoutes(r)
	// 启动服务
	panic(r.Run(":8080"))
}
