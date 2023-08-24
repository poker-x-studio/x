/*
功能：salt
说明：
*/
package xaccount

import (
	"math/rand"
	"strings"
	"time"
)

const (
	ALL_CHARS_SALT = "abcdefghijklmnopqrstuvwxyz0123456789_"
)

// Salt 随机salt
func Salt(length int) string {
	if length <= 0 {
		length = 6
	}
	chars := strings.Split(ALL_CHARS_SALT, "")
	len := len(chars)
	r := rand.New(rand.NewSource(time.Now().UnixMicro()))
	var txt string = ""
	for i := 0; i < length; i++ {
		txt = txt + chars[r.Intn(len)]
	}
	return txt
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
