/*
功能：字符串相关
说明：
*/
package xutils

// 反转字符串
func String_reverse(str string) string {
	runes := []rune(str)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
