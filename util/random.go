package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// 生成一个随机整数，在min和max之间
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// 生成一组随机字符串
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// 随机生成用户名
func RandomOwner() string {
	return RandomString(6)
}

// 随机生成账户现金
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// 随机生成货币种类
func RandomCurrency() string {
	currencies := []string{"RMB", "USD", "EUR"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// 随机生成发送账户
func RandomFromAccountID() int64 {
	return RandomInt(1000000000000000, 9999999999999999)
}

// 随机生成接收账户
func RandomToAccountID() int64 {
	return RandomInt(1000000000000000, 9999999999999999)
}

// 随机生成转账金额
func RandomAmount() int64 {
	return RandomInt(1, 100)
}
