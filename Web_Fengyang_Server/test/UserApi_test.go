package test

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
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	_ "github.com/go-sql-driver/mysql"
)

var router *gin.Engine
var userController controller.IUserController
var jsonData map[string]interface{}
var db *gorm.DB
var testToken []string

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

func cleanTable() {
	db.Exec("DELETE FROM client")
	db.Exec("DELETE FROM admin")
}

func loadTestDataFromFile() {
	// 读取JSON文件内容
	data, err := ioutil.ReadFile("./userTestData.json")
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
	login := jsonData["login"].(map[string]interface{})

	for name := range login {
		t.Run(name, func(t *testing.T) {
			for _, user := range login[name].([]interface{}) {
				userData := user.(map[string]interface{})
				// 构造请求参数
				requestBody := gin.H{
					"phoneNumber": userData["phoneNumber"],
					"password":    userData["password"],
					"userType":    userData["userType"],
				}
				requestBodyJson, _ := json.Marshal(requestBody)

				// 构造请求
				req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(requestBodyJson))
				req.Header.Set("Content-Type", "application/json")

				// 发送请求
				resp := httptest.NewRecorder()
				router.ServeHTTP(resp, req)

				// 断言响应
				passTest := assert.Contains(t, resp.Body.String(), userData["expectResponse"])

				// 读取响应体中的内容

				if passTest && userData["expectResponse"] == "登录成功" {
					// 解码 JSON 格式数据
					var jsonMap map[string]interface{}
					json.Unmarshal(resp.Body.Bytes(), &jsonMap)
					// 提取出 token 字段
					respData := jsonMap["data"].(map[string]interface{})
					testToken = append(testToken, respData["token"].(string))
				}
			}
		})
	}
}

func TestUserController_GetMyInfo(t *testing.T) {
	for _, token := range testToken {
		req, _ := http.NewRequest("GET", "/myinfo", nil)
		bearerToken := "Bearer " + token
		req.Header.Set("Authorization", bearerToken)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		assert.Contains(t, resp.Body.String(), "登录获取信息成功")
	}
}

func TestUserController_List(t *testing.T) {
	list := jsonData["list"].(map[string]interface{})

	for name := range list {
		t.Run(name, func(t *testing.T) {
			for _, queryData := range list[name].([]interface{}) {
				userDataMap := queryData.(map[string]interface{})

				query := "userType=" + userDataMap["userType"].(string)
				query += "&keyword=" + userDataMap["keyword"].(string)
				query += "&pageNum=" + userDataMap["pageNum"].(string)
				query += "&pageSize=" + userDataMap["pageSize"].(string)
				query += "&order=" + userDataMap["order"].(string)
				// 构造请求
				req, _ := http.NewRequest("GET", "/list?"+query, nil)

				// 发送请求
				resp := httptest.NewRecorder()
				router.ServeHTTP(resp, req)

				// 断言响应
				assert.Contains(t, resp.Body.String(), userDataMap["expectResponse"])
			}
		})
	}
}

func getExistUser() []map[string]interface{} {
	register := jsonData["register"].(map[string]interface{})

	var result []map[string]interface{}

	for _, user := range register["success"].([]interface{}) {
		userData := user.(map[string]interface{})

		if userData["status"].(bool) {
			var user model.User
			db.Table(userData["userType"].(string)).Where("phone_number=?", userData["phoneNumber"].(string)).First(&user)

			userData["id"] = strconv.FormatUint(uint64(user.ID), 10)
			result = append(result, userData)
		}
	}

	return result
}

func TestUserController_GetBriefInfo(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		existUser := getExistUser()
		for _, data := range existUser {
			query := "userType=" + data["userType"].(string)
			query += "&id=" + data["id"].(string)
			// 构造请求
			req, _ := http.NewRequest("GET", "/briefinfo?"+query, nil)

			// 发送请求
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			// 断言响应
			assert.Contains(t, resp.Body.String(), "查找成功")
		}
	})

	t.Run("fail", func(t *testing.T) {
		getBriefInfo := jsonData["getBriefInfo"].([]interface{})

		for _, idata := range getBriefInfo {
			data := idata.(map[string]interface{})
			query := "userType=" + data["userType"].(string)
			query += "&id=" + data["id"].(string)
			// 构造请求
			req, _ := http.NewRequest("GET", "/briefinfo?"+query, nil)

			// 发送请求
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			// 断言响应
			assert.Contains(t, resp.Body.String(), data["expectResponse"])
		}
	})
}

func TestUserController_Update(t *testing.T) {
	// t.Run("success", func(t *testing.T) {
	// 	existUser := getExistUser()
	// 	for _, data := range existUser {
	// 		query := "userType=" + data["userType"].(string)
	// 		query += "&id=" + data["id"].(string)
	// 		// 构造请求
	// 		req, _ := http.NewRequest("GET", "/briefinfo?"+query, nil)

	// 		// 发送请求
	// 		resp := httptest.NewRecorder()
	// 		router.ServeHTTP(resp, req)

	// 		// 断言响应
	// 		assert.Contains(t, resp.Body.String(), "查找成功")
	// 	}
	// })

	// t.Run("fail", func(t *testing.T) {
	// 	getBriefInfo := jsonData["getBriefInfo"].([]interface{})

	// 	for _, idata := range getBriefInfo {
	// 		data := idata.(map[string]interface{})
	// 		query := "userType=" + data["userType"].(string)
	// 		query += "&id=" + data["id"].(string)
	// 		// 构造请求
	// 		req, _ := http.NewRequest("GET", "/briefinfo?"+query, nil)

	// 		// 发送请求
	// 		resp := httptest.NewRecorder()
	// 		router.ServeHTTP(resp, req)

	// 		// 断言响应
	// 		assert.Contains(t, resp.Body.String(), data["expectResponse"])
	// 	}
	// })

}

func TestUserController_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		existUser := getExistUser()
		for _, data := range existUser {
			data["id"], _ = strconv.ParseUint(data["id"].(string), 10, 32)
			// 构造请求参数
			requestBody := gin.H{
				"userType": data["userType"],
				"ID":       data["id"],
			}
			requestBodyJson, _ := json.Marshal(requestBody)

			// 构造请求
			req, _ := http.NewRequest("DELETE", "/delete", bytes.NewBuffer(requestBodyJson))
			req.Header.Set("Content-Type", "application/json")

			// 发送请求
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			// 断言响应
			assert.Contains(t, resp.Body.String(), "删除成功")
		}
	})

}
