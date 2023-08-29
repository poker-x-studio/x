/*
功能：时间开关-配置项
说明：
*/
package time_switch

type StartHandler func() //开启函数
type StopHandler func()  //停止函数

type Option func(*TimeSwitch)

//WithStopHandler 停止函数
func WithStopHandler(stop StopHandler) Option {
	return func(t *TimeSwitch) {
		t.stop_handler = stop
	}
}

//WithStarHandler 开启函数
func WithStarHandler(start StartHandler) Option {
	return func(t *TimeSwitch) {
		t.start_handler = start
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
