package routes

import (
	"Web_Fengyang_Server/controller"
	"Web_Fengyang_Server/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {
	// 允许跨域访问
	r.Use(middleware.CORSMiddleware())
	// 注册
	r.POST("/register", controller.Register)
	// 登录
	r.POST("/login", controller.Login)
	// 登录获取用户信息
	r.GET("/user", middleware.AuthMiddleware(), controller.GetInfo)

	return r
}
