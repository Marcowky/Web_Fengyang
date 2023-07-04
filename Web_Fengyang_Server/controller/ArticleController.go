package controller

import (
	"Web_Fengyang_Server/common"
	"Web_Fengyang_Server/model"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type ArticleController struct {
	DB *gorm.DB
}

type IArticleController interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Show(c *gin.Context)
	List(c *gin.Context)
}

func (a ArticleController) Create(c *gin.Context) {
	var articleRequest model.CreateArticleRequest
	// 数据验证
	if err := c.ShouldBindJSON(&articleRequest); err != nil {
		common.Fail(c, 400, nil, "数据错误")
		return
	}
	// 检测文章标题
	var errT error
	articleRequest.Title, errT = common.TextCheck(articleRequest.Title)

	if errT != nil {
		common.Fail(c, 400, nil, "发布失败")
		return
	}

	// 获取登录用户
	user, _ := c.Get("user")
	// 创建文章
	article := model.Article{
		UserId:      user.(model.User).ID,
		CategoryId:  articleRequest.CategoryId,
		Title:       articleRequest.Title,
		Content:     articleRequest.Content,
		HeadImage:   articleRequest.HeadImage,
		ArticleType: articleRequest.ArticleType,
	}
	if err := a.DB.Table(article.ArticleType).Create(&article).Error; err != nil {
		common.Fail(c, 400, nil, "发布失败")
		return
	}

	common.Success(c, gin.H{"id": article.ID}, "发布成功")
}

func (a ArticleController) Update(c *gin.Context) {

	var articleRequest model.CreateArticleRequest
	// 数据验证
	if err := c.ShouldBindJSON(&articleRequest); err != nil {
		common.Fail(c, 400, nil, "数据错误")
		return
	}
	// 获取path中的id
	articleType := c.Query("articleType")
	articleId := c.Query("id")
	// 查找文章
	var article model.Article
	if a.DB.Table(articleType).Where("id = ?", articleId).First(&article).RecordNotFound() {
		common.Fail(c, 400, nil, "文章不存在")
		return
	}

	// 获取登录用户
	user, _ := c.Get("user")
	userId := user.(model.User).ID
	userType := user.(model.User).UserType
	if userId != article.UserId && userType != "admin" {
		common.Fail(c, 400, nil, "登录用户不正确")
		return
	}
	// 更新文章
	if err := a.DB.Table(articleType).Model(&article).Update(articleRequest).Error; err != nil {
		common.Fail(c, 400, nil, "修改失败")
		return
	}
	common.Success(c, nil, "修改成功")
}

func (a ArticleController) Delete(c *gin.Context) {
	// 获取path中的id
	articleType := c.Query("articleType")
	articleId := c.Query("id")
	// 查找文章
	var article model.Article
	if a.DB.Table(articleType).Where("id = ?", articleId).First(&article).RecordNotFound() {
		common.Fail(c, 400, nil, "文章不存在")
		return
	}
	// 获取登录用户
	user, _ := c.Get("user")
	userId := user.(model.User).ID
	userType := user.(model.User).UserType
	if userId != article.UserId && userType != "admin" {
		common.Fail(c, 400, nil, "登录用户不正确")
		return
	}
	// 删除文章
	if err := a.DB.Table(articleType).Delete(&article).Error; err != nil {
		common.Fail(c, 400, nil, "删除失败")
		return
	}

	common.Success(c, nil, "删除成功")
}

func (a ArticleController) Show(c *gin.Context) {
	// 获取path中的id
	articleType := c.Query("articleType")
	articleId := c.Query("id")
	// 查找文章
	var article model.Article
	if a.DB.Table(articleType).Where("id = ?", articleId).First(&article).RecordNotFound() {
		common.Fail(c, 400, nil, "文章不存在")
		return
	}
	// 展示文章详情
	common.Success(c, gin.H{"article": article}, "查找成功")
}

func (a ArticleController) List(c *gin.Context) {
	// 获取关键词、分类、分页参数
	articleType := c.Query("articleType")
	keyword := c.DefaultQuery("keyword", "")
	categoryId := c.DefaultQuery("categoryId", "0")
	order := c.DefaultQuery("order", "created_at desc")
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "100"))
	var query []string
	var args []string
	// 若关键词存在
	if keyword != "" {
		query = append(query, "(title LIKE ? OR content LIKE ?)")
		args = append(args, "%"+keyword+"%")
		args = append(args, "%"+keyword+"%")
	}
	// 若分类存在
	if categoryId != "0" {
		query = append(query, "category_id = ?")
		args = append(args, categoryId)
	}
	// 拼接字符串
	var querystr string
	if len(query) > 0 {
		querystr = strings.Join(query, " AND ")
	}
	// 页面内容
	var article []model.Article
	// 文章总数
	var count int
	// 查询文章
	switch len(args) {
	case 0:
		a.DB.Order(order).Table(articleType).Select("id, user_id, category_id, title, content AS content, head_image, article_type, created_at, updated_at").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&article)
		a.DB.Table(articleType).Model(model.Article{}).Count(&count)
	case 1:
		a.DB.Order(order).Table(articleType).Select("id, user_id, category_id, title, content AS content, head_image, article_type, created_at, updated_at").Where(querystr, args[0]).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&article)
		a.DB.Table(articleType).Model(model.Article{}).Where(querystr, args[0]).Count(&count)
	case 2:
		a.DB.Order(order).Table(articleType).Select("id, user_id, category_id, title, content AS content, head_image, article_type, created_at, updated_at").Where(querystr, args[0], args[1]).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&article)
		a.DB.Table(articleType).Model(model.Article{}).Where(querystr, args[0], args[1]).Count(&count)
	case 3:
		a.DB.Order(order).Table(articleType).Select("id, user_id, category_id, title, content AS content, head_image, article_type, created_at, updated_at").Where(querystr, args[0], args[1], args[2]).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&article)
		a.DB.Table(articleType).Model(model.Article{}).Where(querystr, args[0], args[1], args[2]).Count(&count)
	}
	for i := 0; i < len(article); i++ {
		article[i].Content = regexp.MustCompile("<.*?>").ReplaceAllString(article[i].Content, "")
		runes := []rune(article[i].Content)
		if len(runes) > 80 {
			article[i].Content = string(runes[:80])
		}
	}

	// 展示文章列表
	common.Success(c, gin.H{"article": article, "count": count}, "查找成功")
}

func NewArticleController() IArticleController {
	db := common.GetDB()
	db.Table("infoArticle").AutoMigrate(model.Article{})
	db.Table("attractionArticle").AutoMigrate(model.Article{})
	db.Table("consumptionArticle").AutoMigrate(model.Article{})
	db.Table("blogArticle").AutoMigrate(model.Article{})
	db.Table("hotelArticle").AutoMigrate(model.Article{})
	return ArticleController{DB: db}
}

// // SearchCategory 查询分类
// func (a ArticleController) SearchCategory(c *gin.Context) {
// 	db := common.GetDB()
// 	var categories []model.Category
// 	if err := db.Find(&categories).Error; err != nil {
// 		common.Fail(c, 400, nil, "查找失败")
// 		return
// 	}
// 	common.Success(c, gin.H{"categories": categories}, "查找成功")
// }

// // SearchCategoryName 查询分类名
// func (a ArticleController) SearchCategoryName(c *gin.Context) {
// 	db := common.GetDB()
// 	var category model.Category
// 	// 获取path中的分类id
// 	categoryId := c.Params.ByName("id")
// 	if err := db.Where("id = ?", categoryId).First(&category).Error; err != nil {
// 		common.Fail(c, 400, nil, "分类不存在")
// 		return
// 	}
// 	common.Success(c, gin.H{"categoryName": category.CategoryName}, "查找成功")
// }

// 初始化文章类别
// func InitCategory() {
// 	db := common.GetDB()
// 	db.AutoMigrate(model.Category{})
// 	var category model.Category
// 	if err := db.Where("id = ?", 0).First(&category).Error; err == nil {
// 		return
// 	}
// 	categoryArray := []string{"ALL", "category1", "category2", "category3", "category4"}
// 	for i := 0; i < len(categoryArray); i++ {
// 		newCategory := model.Category{
// 			ID:           &i,
// 			CategoryName: categoryArray[i],
// 		}
// 		db.Create(&newCategory)
// 	}
// }
