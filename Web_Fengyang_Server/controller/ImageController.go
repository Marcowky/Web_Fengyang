package controller

import (
	"Web_Fengyang_Server/common"
	"Web_Fengyang_Server/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type CarouselController struct {
	DB *gorm.DB
}

type ICarouselController interface {
	AddCarousel(c *gin.Context)
	DeleteCarousel(c *gin.Context)
	List(c *gin.Context)
	UpdateCarousel(c *gin.Context)
	CarouselUrlList(c *gin.Context)
}

// UploadImage 上传图像
func UploadImage(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		common.Fail(c, 500, nil, "格式错误")
		return
	}
	filename := header.Filename
	ext := path.Ext(filename)
	// 用上传时间作为文件名
	name := "image_" + time.Now().Format("20060102150405")
	newFilename := name + ext
	out, err := os.Create("static/images/articleImages/" + newFilename)
	if err != nil {
		common.Fail(c, 500, nil, "创建错误")
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		common.Fail(c, 500, nil, "复制错误")
		return
	}

	common.Success(c, gin.H{"filePath": "/articleImages/" + newFilename}, "上传成功")
}

type DeleteImageRequest struct {
	FilePath string `json:"filePath"`
}

// 删除blog头图函数
func DeleteImage(c *gin.Context) {
	var req DeleteImageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// 处理请求体解析错误
		// ...
		return
	}
	imagePath := filepath.Join("static/images", req.FilePath)
	err := os.Remove(imagePath)
	if err != nil {
		common.Fail(c, 500, nil, "删除错误")
		return
	}
	common.Success(c, nil, "删除成功")
}

// RichEditorUploadImage 上传富文本编辑器中的图像
func RichEditorUploadImage(c *gin.Context) {
	fromData, _ := c.MultipartForm()
	files := fromData.File["wangeditor-uploaded-image"]
	var url []string
	for _, file := range files {
		ext := path.Ext(file.Filename)
		name := "image_" + time.Now().Format("20060102150405")
		newFilename := name + ext
		dst := path.Join("./static/images/articleImages", newFilename)
		fileurl := "/articleImages/" + newFilename
		url = append(url, fileurl)
		err := c.SaveUploadedFile(file, dst)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"errno":   1,
				"message": "上传失败",
			})
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"errno": 0,
		"data": gin.H{
			"url": url[0],
		},
	})
}

// UploadCarousel 上传轮播图图像
func UploadCarousel(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		common.Fail(c, 500, nil, "格式错误")
		return
	}
	filename := header.Filename
	ext := path.Ext(filename)
	// 用上传时间作为文件名
	name := "image_" + time.Now().Format("20060102150405")
	//newFilename := category + name + ext
	url := "/carousel/" + name + ext
	out, err := os.Create("static/images" + url)
	if err != nil {
		common.Fail(c, 500, nil, "创建错误")
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		common.Fail(c, 500, nil, "复制错误")
		return
	}

	common.Success(c, gin.H{"filePath": url}, "上传成功")
}

// 添加轮播图
func (a CarouselController) AddCarousel(c *gin.Context) {

	// 获取context中的参数
	var requestCarousel model.Carousel
	c.Bind(&requestCarousel) // 将请求中的JSON数据绑定到User结构体中，方便后续操作
	iorder := requestCarousel.Iorder
	category := requestCarousel.Category
	url := requestCarousel.Url

	// 创建轮播图
	newCarousel := model.Carousel{
		Iorder:   iorder,
		Category: category,
		Url:      url,
	}
	a.DB.Table(category).Create(&newCarousel)

	// 返回成功响应
	common.Success(c, nil, "添加成功")
}

// Update 用于修改轮播图
func (a CarouselController) UpdateCarousel(c *gin.Context) {

	// 获取context中的参数
	var updateCarousel model.Carousel
	c.Bind(&updateCarousel) // 将请求中的JSON数据绑定到User结构体中，方便后续操作
	ID := updateCarousel.ID
	iorder := updateCarousel.Iorder
	category := updateCarousel.Category
	url := updateCarousel.Url
	//println("更新轮播图id")
	//println(ID)
	//println("更新轮播图ioder")
	//println(iorder)
	//println("更新轮播图category" + category)

	// 验证数据
	var carousel model.Carousel
	a.DB.Table(category).Where("id=?", ID).First(&carousel)
	if carousel.ID == 0 {
		common.Fail(c, 422, nil, "轮播图不存在")
		return
	}

	carousel.Iorder = iorder
	carousel.Url = url

	// 执行更新操作
	if err := a.DB.Table(category).Save(&carousel).Error; err != nil {
		common.Fail(c, 422, nil, "修改错误")
		return
	}

	// 返回成功响应
	common.Success(c, nil, "修改成功")
}

// 删除轮播图
func (a CarouselController) DeleteCarousel(c *gin.Context) {
	var deleteCarousel model.Carousel
	c.Bind(&deleteCarousel)
	ID := deleteCarousel.ID
	category := deleteCarousel.Category
	//url := deleteCarousel.Url

	// 从/carousel文件夹中删除
	//imagePath := filepath.Join("static/images", url)
	//err := os.Remove(imagePath)
	//if err != nil {
	//	common.Fail(c, 500, nil, "删除错误")
	//	return
	//}
	// 数据验证
	var carousel model.Carousel
	a.DB.Table(category).Where("id =?", ID).First(&carousel)
	if carousel.ID == 0 {
		common.Fail(c, 422, nil, "轮播图不存在")
		return
	}

	// 删除数据库中轮播图
	if err := a.DB.Table(category).Delete(&carousel).Error; err != nil {
		common.Fail(c, 400, nil, "删除失败")
		return
	}

	common.Success(c, nil, "删除成功")

}

// List 用于获取轮播图列表
func (a CarouselController) List(c *gin.Context) {
	//keyword := c.DefaultQuery("keyword", "")
	category := c.Query("category")
	order := c.DefaultQuery("Iorder", "iorder asc")
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "5"))

	var carousel []model.Carousel

	// 构建查询条件
	query := a.DB.Order(order).Table(category)

	// 获取总数
	var count int
	query.Count(&count)

	// 执行查询
	if err := query.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&carousel).Error; err != nil {
		common.Fail(c, 400, nil, "查询失败")
		return
	}

	common.Success(c, gin.H{"carousel": carousel, "count": count}, "加载成功")
}

// List 用于获取ioder不为0的轮播图的url列表
func (a CarouselController) CarouselUrlList(c *gin.Context) {
	category := c.Query("category")
	var order string
	order = "Iorder asc"
	var carousel []model.Carousel

	// 构建查询条件
	query := a.DB.Order(order).Table(category).Where("iorder > ?", 0).Select("url") // 修改这里选择"url"
	//query := a.DB.Order(order).Table(category).Where("Iorder > ?", 0) // 修改这里选择"url"

	// 获取总数
	var count int
	query.Count(&count)

	// 执行查询
	if err := query.Find(&carousel).Error; err != nil {
		common.Fail(c, 400, nil, "查询失败")
		return
	}

	// 提取URL列表
	var urlList []string
	for _, item := range carousel {
		//println("item.Url= ")
		//println(item.Url)
		urlList = append(urlList, item.Url)
	}
	//println("urlList[0]= ")
	//println(urlList[0])
	common.Success(c, gin.H{"urlList": urlList, "count": count}, "加载成功") // 返回urlList
	//common.Success(c, gin.H{"urlList": carousel, "count": count}, "加载成功")
}

func NewCarouselController() ICarouselController {
	db := common.GetDB()
	db.Table("home").AutoMigrate(model.Carousel{})
	db.Table("consumpAttraction").AutoMigrate(model.Carousel{})
	return CarouselController{DB: db}

}
