/*
功能：文件名
说明：不包括扩展名
*/
package xpath

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	ALL_CHARS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

// 随机文件名
func Rand_filename(length int) string {
	if length < 1 {
		return ""
	}
	chars := strings.Split(ALL_CHARS, "")
	len := len(chars)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	var rchar string = ""
	for i := 0; i < length; i++ {
		rchar = rchar + chars[r.Intn(len)]
	}
	return rchar
}

// 日期文件名
func Date_filename() string {
	//当前时区时间
	time := time.Now()
	return fmt.Sprintf("%04d-%02d-%02d-%02d-%02d-%02d-%d", time.Year(), time.Month(), time.Day(), time.Hour(), time.Minute(), time.Second(), time.Nanosecond())
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
