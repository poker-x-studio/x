/*
功能：随机数
说明：
*/
package xmath

import (
	"math/rand"
	"time"
)

// Intn 随机整数 [0,n)
func Intn(n int) int {
	if n <= 0 {
		panic("")
	}
	rand_seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand_seed.Intn(n)
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
