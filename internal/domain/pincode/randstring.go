package pincode

import (
	"math/rand"
	"time"
)

func randString(length int) string {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 定义包含所有可能字符的字符串
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	// 生成随机字符串
	randomString := make([]byte, length)
	for i := 0; i < length; i++ {
		randomString[i] = chars[rand.Intn(len(chars))]
	}

	return string(randomString)
}
