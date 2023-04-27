package controller

import (
	"Web_Fengyang_Server/common"
	"Web_Fengyang_Server/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// 注册
func Register(c *gin.Context) {
	// 连接当前数据库
	db := common.GetDB()

	// 获取context中的参数
	var requestUser model.User
	c.Bind(&requestUser) // 将请求中的JSON数据绑定到User结构体中，方便后续操作
	userName := requestUser.UserName
	phoneNumber := requestUser.PhoneNumber
	password := requestUser.Password

	// 验证数据
	var user model.User
	db.Where("phone_number=?", phoneNumber).First(&user)
	if user.ID != 0 { // 若已存在电话号码，则返回422，且返回提示信息
		c.JSON(http.StatusOK, gin.H{
			"code": 422,
			"msg":  "电话号码已注册",
		})
		return
	}

	// 密码加密
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	// 创建用户
	newUser := model.User{
		UserName:    userName,
		PhoneNumber: phoneNumber,
		Password:    string(hashedPassword),
	}
	db.Create(&newUser)

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功",
	})

}

// Login 登录
func Login(c *gin.Context) {
	db := common.GetDB()

	// 获取参数
	var requestUser model.User
	c.Bind(&requestUser)
	phoneNumber := requestUser.PhoneNumber
	password := requestUser.Password

	// 数据验证
	var user model.User
	db.Where("phone_number =?", phoneNumber).First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 422,
			"msg":  "用户不存在",
		})
		return
	}

	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 422,
			"msg":  "密码错误",
		})
		return
	}

	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "系统异常",
		})
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"token": token},
		"msg":  "登录成功",
	})
}

// GetInfo 登录后获取信息，用于测试中间件
func GetInfo(c *gin.Context) {
	// 获取上下文中的用户信息
	user, _ := c.Get("user")

	// 返回用户信息
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"id": user.(model.User).ID},
		"msg":  "登录获取信息成功",
	})
}
