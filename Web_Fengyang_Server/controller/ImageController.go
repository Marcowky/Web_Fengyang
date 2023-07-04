package controller

import (
	"Web_Fengyang_Server/common"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

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
