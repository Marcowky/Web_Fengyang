package main

import (
	"Web_Fengyang_Server/common"
	"Web_Fengyang_Server/controller"
	"Web_Fengyang_Server/routes"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 获取初始化的数据库
	db := common.InitDB()
	// 初始化文章分类
	controller.InitCategory()
	// 延迟关闭数据库
	defer db.Close()
	// 创建路由引擎
	r := gin.Default()
	// 配置静态文件路径
	r.StaticFS("/images", http.Dir("./static/images"))
	// 启动路由
	routes.CollectRoutes(r)
	// 启动服务
	panic(r.Run(":8080"))
}
