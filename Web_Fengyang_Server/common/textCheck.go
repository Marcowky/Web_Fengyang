package common

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

//func HandleFilterText(c *gin.Context) {
//	var requestData struct {
//		ArticleText string `json:"article_text"`
//	}
//
//	if err := c.BindJSON(&requestData); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
//		return
//	}
//
//	filteredText, err := TextCheck(requestData.ArticleText)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error filtering text"})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"filteredText": filteredText})
//}

func TextCheck(text string) (string, error) {

	uri := "https://eolink.o.apispace.com/text-filters/api/v1/forward/text_filter/"

	payload := url.Values{}
	payload.Set("text", text)
	payload.Set("symbol", "*")
	req, _ := http.NewRequest("POST", uri, strings.NewReader(payload.Encode()))

	req.Header.Add("X-APISpace-Token", "w7aeug15gzbtubqujsxokbh323ury6he")
	req.Header.Add("Authorization-Type", "apikey")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	//println(res)
	defer res.Body.Close()
	//body, err := ioutil.ReadAll(res.Body)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
		//println("TextCheck error")
	}
	//println("TextCheck" + string(body))

	var result map[string]interface{}
	errjs := json.Unmarshal(body, &result)

	if errjs != nil {
		panic(errjs)
	}

	// 获取"data"字段的值
	dataValue := result["data"].(string)

	fmt.Println(dataValue)
	return dataValue, nil
}
