package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Post(url string, data interface{}, headers map[string]string) (res string, err error) {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		return "", err

	}
	// 添加请求头
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	req.Header.Set("accept", "application/json")
	client := &http.Client{}
	client.Timeout = 15 * time.Second
	resp, err := client.Do(req)
	if err != nil {
		return "", err

	}
	defer func() {
		_ = resp.Body.Close()
	}()
	fmt.Println("status", resp.StatusCode)
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body), nil
}

func Get(url string, qs map[string]string, headers map[string]string) (res string, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	q := req.URL.Query()
	// 添加请求参数
	for k, v := range qs {
		q.Add(k, v)
	}
	// 添加请求头
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	client := &http.Client{}
	client.Timeout = 15 * time.Second
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	fmt.Println("statusCode", req.URL, resp.StatusCode)
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), nil
}
