package crypto

import "testing"
import "reflect"

// TestGetMD5 测试md5生成
func TestGetMD5(t *testing.T) {
	// single params
	s := "Hello,Golang!"
	md := "9cea3f44c8955043a21997961b5f60a2"

	if !reflect.DeepEqual(GetMD5(s), md) {
		t.Errorf("o MD5Test failt,%s md5 should be %s ,but %s", s, md, GetMD5(s))
	}
	// mutiple params
	d := "Hello,Node"
	mdd := "b5a19c344179a1943449630d98515236"

	if !reflect.DeepEqual(GetMD5(s, d), mdd) {
		t.Errorf("o MD5Test failt,%s md5 should be %s ,but %s", s, md, GetMD5(s))
	}
}
