package util

import (
	"math/rand"
	"strings"
	"time"
)

var (
	rnd      *rand.Rand
	alphabet = "abcdefghijklmnopqrstuvwxyz"
)

func init() {
	source := rand.NewSource(time.Now().UnixNano())
	rnd = rand.New(source)
}

func RandomInt(min, max int64) int64 {
	return min + rnd.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rnd.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(1, 1000)
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "KZT"}
	n := len(currencies)
	return currencies[rnd.Intn(n)]
}
