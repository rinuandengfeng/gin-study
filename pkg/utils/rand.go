package utils

import (
	"math/rand"
	"time"
	"unsafe"
)

const (
	letterInt    = "0123456789"
	letter       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdBits = 6
	letterIdMask = 1<<letterIdBits - 1
	letterIdMax  = 63 / letterIdBits
)

var randSource = rand.NewSource(time.Now().UnixNano())

// RandInt 返回在区间中的随机数.
//
//nolint:gosec
func RandInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

// RandNum 返回指定长度的数值随机字符串.
func RandNum(length int) string {
	return randValue(length, letterInt)
}

// RandStr 返回指定长度的随机数值字母组合字符串.
func RandStr(length int) string {
	return randValue(length, letter)
}

func randValue(n int, letters string) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdMax letters!
	for i, cache, remain := n-1, randSource.Int63(), letterIdMax; i >= 0; {
		if remain == 0 {
			cache, remain = randSource.Int63(), letterIdMax
		}

		if idx := int(cache & letterIdMask); idx < len(letters) {
			b[i] = letters[idx]
			i--
		}

		cache >>= letterIdBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}
