package main

import (
	"Web_Fengyang_Server/common"
	"Web_Fengyang_Server/controller"
	"Web_Fengyang_Server/model"
	"Web_Fengyang_Server/routes/middleware"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/jinzhu/gorm"
)

var router *gin.Engine
var userController controller.IUserController
var jsonData map[string]interface{}
var db *gorm.DB

func init() {
	loadTestDataFromFile()

	db = common.InitDB("test")
	db.Table("client").AutoMigrate(model.User{})
	db.Table("admin").AutoMigrate(model.User{})

	userController = controller.NewUserController()
	router = gin.Default()
	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)
	router.PUT("/update", userController.Update)
	router.DELETE("/delete", userController.Delete)
	router.GET("/list", userController.List)
	router.GET("/myinfo", middleware.AuthMiddleware(), userController.GetMyInfo)
	router.GET("/briefinfo", userController.GetBriefInfo)
}

func TestMain(m *testing.M) {
	result := m.Run()
	cleanTable()
	os.Exit(result)
}

func cleanTable(){
	db.Exec("DELETE FROM client")
	db.Exec("DELETE FROM admin")
}

func loadTestDataFromFile() {
	// 读取JSON文件内容
	data, err := ioutil.ReadFile("./test/userTestData.json")
	if err != nil {
		fmt.Println("无法读取文件:", err)
		return
	}

	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		fmt.Println("无法解析JSON数据:", err)
		return
	}
}

func TestUserController_Register(t *testing.T) {
	register := jsonData["register"].(map[string]interface{})

	for name := range register {
		t.Run(name, func(t *testing.T) {
			for _, user := range register[name].([]interface{}) {
				userData := user.(map[string]interface{})
				// 构造请求参数
				requestBody := gin.H{
					"userName":    userData["userName"],
					"phoneNumber": userData["phoneNumber"],
					"password":    userData["password"],
					"userType":    userData["userType"],
					"status":      userData["status"],
				}
				requestBodyJson, _ := json.Marshal(requestBody)

				// 构造请求
				req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(requestBodyJson))
				req.Header.Set("Content-Type", "application/json")

				// 发送请求
				resp := httptest.NewRecorder()
				router.ServeHTTP(resp, req)

				// 断言响应
				assert.Contains(t, resp.Body.String(), userData["expectResponse"])
			}
		})
	}
}

func TestUserController_Login(t *testing.T) {
	// ...
}

func TestUserController_Update(t *testing.T) {
	// ...
}

func TestUserController_Delete(t *testing.T) {
	// ...
}

func TestUserController_List(t *testing.T) {
	// ...
}

func TestUserController_GetMyInfo(t *testing.T) {
	// ...
}

func TestUserController_GetBriefInfo(t *testing.T) {
	// ...
}
