package utils

import (
	"math/rand"
	"time"
)

func GetRandomString(length int) string {
	str := "zxcvbnmlkjhgfdsaqwertyuiopQWERTYUIOPASDFGHJKLZXCVBNM1234567890"
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var sb []byte
	for i := 0; i < length; i++ {
		number := r.Int() % 62
		sb = append(sb, str[number])
	}
	return string(sb)
}
