/*
功能：
说明：
*/
package xdebug

import "github.com/sirupsen/logrus"

func init() {
	startup_output()
	Init_logrus()
	//Get_sys_info()
}

// Init_logrus logrus日志初始化设置
func Init_logrus() {
	if Is_debug() {
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
