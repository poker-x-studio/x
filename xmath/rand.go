/*
功能：随机数
说明：
*/
package xmath

import (
	"fmt"
	"math/rand"
	"time"
)

// Intn 随机整数 [0,n)
func Intn(n int) (int, error) {
	if n <= 0 {
		return 0, fmt.Errorf("invalid input: n must be positive, got %d", n)
	}
	rand_seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand_seed.Intn(n), nil
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
