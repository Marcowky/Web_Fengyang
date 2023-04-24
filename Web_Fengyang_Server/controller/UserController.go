package controller

import (
	"Web_Fengyang_Server/common"
	"Web_Fengyang_Server/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

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
