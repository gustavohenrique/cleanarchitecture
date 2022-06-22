package random

import (
	crand "crypto/rand"
	"encoding/base64"
	"math/rand"
	"strconv"
	"time"
	"unsafe"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits
)

var src = rand.NewSource(time.Now().UnixNano())

func Strings(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}

func SixNumbers() string {
	rand.Seed(time.Now().UnixNano())
	code := randomInt(100000, 999999)
	return strconv.Itoa(code)
}

func EightNumbers() string {
	rand.Seed(time.Now().UnixNano())
	code := randomInt(10000000, 99999999)
	return strconv.Itoa(code)
}

func Token(size int) string {
	b := make([]byte, size)
	crand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}
