/*
功能：goroutine信息
说明：
*/
package xdebug

import "github.com/petermattis/goid"

// 得到goroutine的id
func Go_id() int64 {
	return goid.Get()
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
