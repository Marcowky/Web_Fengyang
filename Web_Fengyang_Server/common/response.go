package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	c.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}

// Success 成功
func Success(c *gin.Context, data gin.H, msg string) {
	Response(c, http.StatusOK, 200, data, msg)
}

// Fail 失败
func Fail(c *gin.Context, code int, data gin.H, msg string) {
	if code == 400 { // 请求错误
		Response(c, http.StatusOK, 400, data, msg)
	} else if code == 422 { // 内容无效
		Response(c, http.StatusOK, 422, data, msg)
	} else if code == 500 { // ServerError 服务器错误
		Response(c, http.StatusInternalServerError, 500, data, msg)
	} else if code == 401 { // Unauthorized 权限不足
		Response(c, http.StatusUnauthorized, 401, data, msg)
	}
}
