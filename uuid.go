package uuid

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func GenerateOrderedUUID() string {
	now := time.Now().UTC().UnixNano()
	timeStamp := fmt.Sprintf("%x\n", now)
	var arr []string
	arr = append(arr, timeStamp[:8])
	arr = append(arr, timeStamp[8:12])
	arr = append(arr, timeStamp[12:16])
	arr = append(arr, randStringRunes(4))
	arr = append(arr, randStringRunes(12))
	return strings.Join(arr, "-")
}

func randStringRunes(n int) string {
	var letterRunes = []rune("abcdef0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
