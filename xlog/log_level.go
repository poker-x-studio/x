/*
功能：日志等级
说明：
*/
package xlog

import (
	"github.com/poker-x-studio/x/xdebug"

	"github.com/sirupsen/logrus"
)

// set_level 设置日志等级
func set_level() {
	logrus.SetLevel(get_level())
}

// get_level 获取日志等级
func get_level() logrus.Level {
	//所有debug标志统一在xdebug中设置
	if xdebug.Is_debug() {
		return logrus.TraceLevel
	}
	return logrus.InfoLevel
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
