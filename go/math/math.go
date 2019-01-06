package math

import "math/rand"

// GetRandomInt 生成max~min 随机整数
func GetRandomInt(max int, min int) int {
	return rand.Intn(max-min) + min
}

// GetRandomNumberString 生成随机长度的整数字符串
func GetRandomNumberString(l int) string {
	s := "1234567890"
	b := make([]byte, l)
	for i := range b {
		b[i] = s[rand.Intn(len(s))]
	}
	return string(b)
}
