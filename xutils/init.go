/*
功能：
说明：
*/
package xutils

import (
	"github.com/poker-x-studio/x/xdebug"

	"github.com/sirupsen/logrus"
)

func init() {
	if xdebug.Is_debug() {
		logrus.SetLevel(logrus.TraceLevel)
		logrus.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05", //时间格式
			FullTimestamp:   true,
		})
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
