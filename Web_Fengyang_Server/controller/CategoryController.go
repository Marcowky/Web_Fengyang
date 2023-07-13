package controller

import (
	"Web_Fengyang_Server/common"
	"Web_Fengyang_Server/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CategoryController struct {
	DB *gorm.DB
}

type ICategoryController interface {
	CategoryList(c *gin.Context)
}

func NewCategoryController() ICategoryController {
	db := common.GetDB()
	db.Table("category").AutoMigrate(model.Category{})
	return CategoryController{DB: db}
}

func (a CategoryController) CategoryList(c *gin.Context) {
	// 页面内容
	var categorys []model.Category
	// 查询文章
	a.DB.Table("category").Find(&categorys)
	// 展示文章列表
	common.Success(c, gin.H{"categorys": categorys}, "获取成功")
}