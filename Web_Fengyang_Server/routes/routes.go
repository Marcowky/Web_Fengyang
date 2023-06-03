package routes

import (
	"Web_Fengyang_Server/controller"
	"Web_Fengyang_Server/middleware"

	"github.com/gin-gonic/gin"
)

/* routes/routes.go */
func CollectRoutes(r *gin.Engine) *gin.Engine {
	// 允许跨域访问
	r.Use(middleware.CORSMiddleware())

	// 图像操作
	imageRoutes := r.Group("/image")
	{
		imageRoutes.POST("upload", controller.UploadImage)                              // 通用图像上传
		imageRoutes.POST("upload/rich_editor_upload", controller.RichEditorUploadImage) // 富文本编辑器内部图像上传
		imageRoutes.POST("delete", controller.DeleteImage)                              // blog头图删除
	}

	// 用户操作
	userController := controller.NewUserController()
	userRoutes := r.Group("/user")
	{
		userRoutes.POST("register", userController.Register)                          // 注册
		userRoutes.POST("login", userController.Login)                                // 登录
		userRoutes.GET("info", middleware.AuthMiddleware(), userController.GetMyInfo) // 获取当前用户信息
		userRoutes.GET("briefInfo", userController.GetBriefInfo)                      // 获取文章作者简要信息
		userRoutes.PUT("delete", userController.Delete)                            // 
		userRoutes.PUT("update", userController.Update)                               // 删除
		userRoutes.GET("list", userController.List)                                   // 获取用户列表
	}

	//文章操作
	articleController := controller.NewArticleController()
	articleRoutes := r.Group("/article")
	{
		articleRoutes.POST("create", middleware.AuthMiddleware(), articleController.Create)   // 发布文章
		articleRoutes.PUT("update", middleware.AuthMiddleware(), articleController.Update)    // 修改文章
		articleRoutes.DELETE("delete", middleware.AuthMiddleware(), articleController.Delete) // 删除文章
		articleRoutes.GET("detail", articleController.Show)                                   // 查看文章
		articleRoutes.GET("list", articleController.List)                                     // 显示文章列表
		// articleRoutes.GET("category", articleController.SearchCategory)                    // 查询分类
		// articleRoutes.GET("category/:id", articleController.SearchCategoryName)            // 查询分类名

	}

	return r
}
