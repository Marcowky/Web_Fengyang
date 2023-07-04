package middleware

import (
	"Web_Fengyang_Server/common"
	"Web_Fengyang_Server/model"
	"strings"

	"github.com/gin-gonic/gin"
)

// 验证token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取authorization header
		tokenString := c.Request.Header.Get("Authorization")

		// token为空
		if tokenString == "" {
			common.Fail(c, 401, nil, "未登录")
			c.Abort()
			return
		}

		// 非法token
		if tokenString == "" || len(tokenString) < 7 || !strings.HasPrefix(tokenString, "Bearer") {
			common.Fail(c, 401, nil, "非法权限")
			c.Abort()
			return
		}

		// 提取token的有效部分
		tokenString = tokenString[7:]
		// 解析token
		token, claims, err := common.ParseToken(tokenString)

		// 非法token
		if err != nil || !token.Valid {
			common.Fail(c, 401, nil, "非法权限")
			c.Abort()
			return
		}

		// token正常，获取claims中的userId
		userId := claims.UserId
		userType := claims.UserType
		DB := common.GetDB()
		var user model.User
		DB.Table(userType).Where("id =?", userId).First(&user)

		// 将用户信息写入上下文便于读取
		c.Set("user", user)
		c.Next()
	}
}
