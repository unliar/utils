package http

import (
	"io/ioutil"
	"net/http"
)

// Get 用于发送http get请求
func Get(s string) (string, error) {
	resp, err := http.Get(s)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil

}
