package controller

import (
	"Web_Fengyang_Server/common"
	"Web_Fengyang_Server/model"

	"github.com/gin-gonic/gin"
)

/* controller/CategoryController.go */
// SearchCategory 查询分类
func SearchCategory(c *gin.Context) {
	db := common.GetDB()
	var categories []model.Category
	if err := db.Find(&categories).Error; err != nil {
		common.Fail(c, 400, nil, "查找失败")
		return
	}
	common.Success(c , gin.H{"categories": categories}, "查找成功")
 }
 
 // SearchCategoryName 查询分类名
 func SearchCategoryName(c *gin.Context) {
	db := common.GetDB()
	var category model.Category
	// 获取path中的分类id
	categoryId := c.Params.ByName("id")
	if err := db.Where("id = ?", categoryId).First(&category).Error; err != nil {
	   common.Fail(c, 400, nil, "分类不存在")
	   return
	}
	common.Success(c, gin.H{"categoryName": category.CategoryName}, "查找成功")
 }
 