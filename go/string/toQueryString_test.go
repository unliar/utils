package string

import (
	"strings"
	"testing"
)

func TestM_ToQueryString(t *testing.T) {
	m := M{
		"code": "1",
		"data": "ddd",
		"em":   "",
	}
	str := m.ToQueryString(false, false)

	if str == "" || !strings.Contains(str, "em") {
		t.Error("TestM_ToQueryString fail")
	}
}
