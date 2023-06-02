package controller

import (
	"Web_Fengyang_Server/common"
	"Web_Fengyang_Server/model"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	DB *gorm.DB
}

type IUserController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	List(c *gin.Context)
	GetMyInfo(c *gin.Context)
	GetBriefInfo(c *gin.Context)
}

// 注册
func (a UserController) Register(c *gin.Context) {

	// 获取context中的参数
	var requestUser model.User
	c.Bind(&requestUser) // 将请求中的JSON数据绑定到User结构体中，方便后续操作
	userName := requestUser.UserName
	phoneNumber := requestUser.PhoneNumber
	password := requestUser.Password
	userType := requestUser.UserType

	// 验证数据
	var user model.User
	a.DB.Table(userType).Where("phone_number=?", phoneNumber).First(&user)
	if user.ID != 0 { // 若已存在电话号码，则返回422，且返回提示信息
		common.Fail(c, 422, nil, "电话号码已注册")
		return
	}

	// 密码加密
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	status := true
	if userType == "admin" {
		status = false
	}
	// 创建用户
	newUser := model.User{
		UserName:    userName,
		PhoneNumber: phoneNumber,
		Password:    string(hashedPassword),
		UserType:    userType,
		Status:      status,
	}
	a.DB.Table(userType).Create(&newUser)

	// 返回成功响应
	common.Success(c, nil, "注册成功")
}

// Login 登录
func (a UserController) Login(c *gin.Context) {

	// 获取参数
	var requestUser model.User
	c.Bind(&requestUser)
	phoneNumber := requestUser.PhoneNumber
	password := requestUser.Password
	userType := requestUser.UserType

	// 数据验证
	var user model.User
	a.DB.Table(userType).Where("phone_number =?", phoneNumber).First(&user)
	if user.ID == 0 {
		common.Fail(c, 422, nil, "用户不存在")
		return
	}

	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		common.Fail(c, 422, nil, "密码错误")
		return
	}

	if !user.Status {
		common.Fail(c, 422, nil, "用户被禁用或正在等待审核")
		return
	}

	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		common.Fail(c, 500, nil, "系统异常")
		return
	}

	// 返回结果
	common.Success(c, gin.H{"token": token}, "登录成功")
}

// Update 用于修改用户信息
func (a UserController) Update(c *gin.Context) {

	// 获取context中的参数
	var updateUser model.User
	c.Bind(&updateUser) // 将请求中的JSON数据绑定到User结构体中，方便后续操作
	ID := updateUser.ID
	userName := updateUser.UserName
	phoneNumber := updateUser.PhoneNumber
	// password := updateUser.Password
	userType := updateUser.UserType
	status := updateUser.Status

	fmt.Println(updateUser.Status)
	fmt.Println(updateUser.Password)

	// 验证数据
	var user model.User
	a.DB.Table(userType).Where("phone_number=?", phoneNumber).First(&user)
	if !(user.ID == ID || user.ID == 0) { // 若已存在电话号码，则返回422，且返回提示信息
		common.Fail(c, 422, nil, "电话号码已存在")
		return
	}

	a.DB.Table(userType).Where("id =?", ID).First(&user)
	if user.ID == 0 {
		common.Fail(c, 422, nil, "用户不存在")
		return
	}

	// // 密码加密
	// hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user.UserName = userName
	user.PhoneNumber = phoneNumber
	// user.Password = string(hashedPassword)
	user.Status = status

	// 执行更新操作
	if err := a.DB.Table(userType).Save(&user).Error; err != nil {
		common.Fail(c, 422, nil, "修改错误")
		return
	}

	// 返回成功响应
	common.Success(c, nil, "修改成功")
}

// Delete 用于删除用户信息
func (a UserController) Delete(c *gin.Context) {

	// 获取参数
	var deleteUser model.User
	c.Bind(&deleteUser)
	ID := deleteUser.ID
	userType := deleteUser.UserType

	// 数据验证
	var user model.User
	a.DB.Table(userType).Where("id =?", ID).First(&user)
	if user.ID == 0 {
		common.Fail(c, 422, nil, "用户不存在")
		return
	}

	// 删除文章
	if err := a.DB.Table(userType).Delete(&user).Error; err != nil {
		common.Fail(c, 400, nil, "删除失败")
		return
	}

	// 返回结果
	common.Success(c, nil, "删除成功")
}

// List 用于获取用户列表
func (a UserController) List(c *gin.Context) {
	keyword := c.Query("keyword")
	userType := c.Query("userType")
	var user []model.User
	// 构建查询条件
	query := a.DB.Table(userType).Where("user_name LIKE ? OR phone_number LIKE ? OR id = ?", "%"+keyword+"%", "%"+keyword+"%", keyword)

	// 执行查询
	if err := query.Find(&user).Error; err != nil {
		common.Fail(c, 400, nil, "查询失败")
		return
	}

	common.Success(c, gin.H{"user": user}, "查找成功")
}

// GetInfo 登录后获取信息，用于测试中间件
func (a UserController) GetMyInfo(c *gin.Context) {
	// 获取上下文中的用户信息
	user, _ := c.Get("user")

	// 返回用户信息
	common.Success(c, gin.H{"id": user.(model.User).ID, "userName": user.(model.User).UserName, "userType": user.(model.User).UserType}, "登录获取信息成功")
}

// GetBriefInfo 通过id获取简要信息
func (a UserController) GetBriefInfo(c *gin.Context) {
	userType := c.Query("userType")
	// 获取path中的userId
	userId := c.Query("id")
	// 判断用户身份
	// user, _ := c.Get("user")
	//var self bool
	var curUser model.User
	a.DB.Table(userType).Where("id =?", userId).First(&curUser)
	if curUser.ID == 0 {
		a.DB.Table("admin").Where("id =?", userId).First(&curUser)
		if curUser.ID == 0 {
			common.Fail(c, 400, nil, "用户不存在")
			return
		}
	}
	// 返回用户简要信息
	common.Success(c, gin.H{"id": curUser.ID, "userName": curUser.UserName, "userType": curUser.UserType}, "查找成功")
}

func NewUserController() IUserController {
	db := common.GetDB()
	db.Table("client").AutoMigrate(model.User{})
	db.Table("admin").AutoMigrate(model.User{})
	return UserController{DB: db}
}
