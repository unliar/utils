package math

import (
	"testing"
)

func TestGetRandomInt(t *testing.T) {
	max, min := 10, 1
	r := GetRandomInt(max, min)
	if r > max || r < min {
		t.Error("out of range", r)
	}
}
