package http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

func Post(url string, data interface{}, headers map[string]string) (res []byte, err error) {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		return nil, err

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
		return nil, err

	}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, _ := ioutil.ReadAll(resp.Body)

	return body, nil
}

func Get(url string, qs map[string]string, headers map[string]string) (res []byte, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}
