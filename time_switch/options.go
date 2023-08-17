/*
功能：时间开关-配置项
说明：
*/
package time_switch

type HandlerStart func() //开启函数
type HandlerStop func()  //停止函数

type Option func(*TimeSwitch)

func WithStopHandler(stop HandlerStop) Option {
	return func(t *TimeSwitch) {
		t.handler_stop = stop
	}
}

func WithStarHandler(start HandlerStart) Option {
	return func(t *TimeSwitch) {
		t.handler_start = start
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
