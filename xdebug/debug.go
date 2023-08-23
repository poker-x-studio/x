/*
功能：调试相关
说明：
所有debug标志统一在xdebug中设置
*/
package xdebug

import "runtime"

var (
	is_debug = true //默认为调试模式
)

const (
	FUNC_ENTER = "进入函数"
	FUNC_EXIT  = "退出函数"
)

//Set_release 设置发布模式
func Set_release() {
	is_debug = false
}

// Is_debug 是否调试模式
func Is_debug() bool {
	return is_debug
}

// Funcname 调用函数名
func Funcname() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
