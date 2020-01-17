package math

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// GetRandomInt 生成max~min 随机整数
func GetRandomInt(max int, min int) int {
	return rand.Intn(max-min) + min
}

// GetRandomNumberString 生成随机长度的整数字符串
func GetRandomNumberString(length int) string {
	s := "1234567890"
	b := make([]byte, length)
	for i := range b {
		b[i] = s[rand.Intn(len(s))]
	}
	return string(b)
}

// 随机长度的字符串
func GetRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// 获取浮点数精度
func FloatDecimal(f float64,n int)(r float64,err error)  {
	nx:= strconv.Itoa(n)
	fm:="%."+nx+"f"
	str:= fmt.Sprintf(fm,f)
	return strconv.ParseFloat(str,10)
}