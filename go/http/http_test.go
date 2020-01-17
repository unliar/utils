package http

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	s, err := Get("https://baidu.com", nil, nil)
	fmt.Println(string(s))
	if err != nil {
		t.Error("TestGet请求失败")
	}
	if s == nil {
		t.Error("TestPost请求失败")
	}
}

func TestPost(t *testing.T) {
	s, err := Post("https://baidu.com", nil, nil)
	fmt.Println(string(s))
	if err != nil {
		t.Error("TestPost请求失败")
	}
	if s == nil {
		t.Error("TestPost请求失败")
	}
}
