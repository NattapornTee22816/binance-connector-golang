package lib

import (
	"bytes"
	"math/rand"
	"time"
)

func GetTimestamp(offset int64) int64 {
	return time.Now().UnixMilli() + offset
}

func ConvertIntToTime(i int64, offset int64) time.Time {
	return time.UnixMilli(i + offset)
}

func BytesToJsonArray(b []byte) []byte {
	if !bytes.HasPrefix(b, []byte("[")) {
		b = append([]byte("["), b...)
		b = append(b, []byte("]")...)
	}
	return b
}

func RandomInt() int64 {
	rand.Seed(time.Now().UnixNano())
	min := 1000
	max := 9999

	return int64(rand.Intn(max-min+1) + min)
}
