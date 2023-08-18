/*
功能：
说明：
*/
package xredis

import (
	"github.com/sirupsen/logrus"
)

const (
	TAG = "db_redis"
)

// 日志
func Log() *logrus.Entry {
	logrus_entry := logrus.WithFields(logrus.Fields{
		"tag": TAG,
	})
	return logrus_entry
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
