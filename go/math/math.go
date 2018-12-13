package math

import "math/rand"

// GetRandomInt 生成max~min 随机整数
func GetRandomInt(max int, min int) int {
	return rand.Intn(max-min) + min
}
