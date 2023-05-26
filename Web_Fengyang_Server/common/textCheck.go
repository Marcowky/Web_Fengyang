package common

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func TextCheck(text string) (string, error) {
	//[]byte
	uri := "https://eolink.o.apispace.com/text-filters/api/v1/forward/text_filter/"

	payload := url.Values{}
	payload.Set("text", text)
	payload.Set("symbol", "{过滤}")
	req, _ := http.NewRequest("POST", uri, strings.NewReader(payload.Encode()))

	req.Header.Add("X-APISpace-Token", "i4gdsvoxjj4t7edn3qpdzfvfp02maroc")
	req.Header.Add("Authorization-Type", "apikey")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
