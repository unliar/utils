package http

import (
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	s, err := Get("https://baidu.com", nil, nil)
	if err != nil {
		if !reflect.DeepEqual(s, "") {
			t.Error("when error must be empty string", s)
		}
	} else {
		if len(s) == 0 {
			t.Error("must have someting")
		}
	}
}
