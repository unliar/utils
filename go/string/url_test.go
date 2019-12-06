package string

import (
	"testing"
)

func TestIsValidUrl(t *testing.T) {
	u := "https://happysoonr.com"
	x := "dadadadadad.cc.cc..cc"
	if r := IsValidUrl(u); r == false {
		t.Errorf("%s is a url", u)
	}
	if r := IsValidUrl(x); r == true {
		t.Errorf("%s is no a url", x)
	}
}
