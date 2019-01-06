package math

import (
	"fmt"
	"testing"
)

func TestGetRandomInt(t *testing.T) {
	max, min := 10, 1
	r := GetRandomInt(max, min)
	if r > max || r < min {
		t.Error("out of range", r)
	}
	fmt.Println("TestGetRandomInt", r)
}

func TestGetRandomNumberString(t *testing.T) {
	l := 6
	r := GetRandomNumberString(l)
	if len(r) != l {
		t.Error("length should be ", l)
	}
	fmt.Println("TestGetRandomNumberString", r)
}
