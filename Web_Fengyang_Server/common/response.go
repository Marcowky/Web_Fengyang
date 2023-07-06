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
	switch code {
	case 400:
		Response(c, http.StatusBadRequest, 400, data, msg) // 请求错误
	case 422:
		Response(c, http.StatusUnprocessableEntity, 422, data, msg) // 内容无效
	case 500:
		Response(c, http.StatusInternalServerError, 500, data, msg) // ServerError 服务器错误
	case 401:
		Response(c, http.StatusUnauthorized, 401, data, msg) // Unauthorized 权限不足
	default:
		Response(c, http.StatusBadRequest, 400, data, msg) // 请求错误
	}
}
