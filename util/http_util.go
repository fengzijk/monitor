package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//发送post请求
func HttpPost(url, postData string) (code int, str string) {

	resp, err := http.Post(url,
		"application/json",
		strings.NewReader(postData))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// handle error
		fmt.Println("http 请求出现错误", err)
		return -1, ""
	}

	fmt.Println(string(body))

	return resp.StatusCode, string(body)
}
