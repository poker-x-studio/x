/*
功能：mysql操作对应的文件日志
说明：
*/
package mysql_logger

import (
	"fmt"
	"strings"
	"x/xlog"
	"x/xlog/async_hook"
	"x/xlog/local_file_hook"

	"github.com/sirupsen/logrus"
)

const (
	mysql_key = "mysql"
)

// fileLogger ...
type fileLogger struct {
	key string
}

func newFileLogger(dir, filename string) (*fileLogger, func(), error) {
	key := fmt.Sprintf("key_%v", filename)

	fileHook := local_file_hook.NewLocalFileHook(dir, filename, xlog.AllLevels)
	asyncHook := async_hook.NewAsyncHookWithHook(fileHook, func(entry *logrus.Entry) bool {

		if val, bfind := entry.Data[key]; bfind {
			if strval, ok := val.(string); ok {
				if strval == mysql_key {
					ldata := make(logrus.Fields, 6)
					for k, v := range entry.Data {
						if k != key {
							ldata[k] = v
						}
					}
					entry.Data = ldata
					return true
				}
			}
		}

		return false
	})

	xlog.StandardLogger().AddHook(asyncHook)

	return &fileLogger{key: key}, func() {
		asyncHook.Close()
		fileHook.Close()
	}, nil
}

// Print ...
func (f *fileLogger) Printf(msg string, data ...interface{}) {
	len := len(data)
	for i := 0; i < len; i++ {
		v, k := data[i].(string)
		if k { //字符串前/后增加一个空格,方便查看
			if i == 0 {
				data[i] = fmt.Sprintf("%s ", v)
				continue
			} else if i == (len - 1) {
				data[i] = fmt.Sprintf(" %s", v)
			} else {
				data[i] = fmt.Sprintf(" %s ", v)
			}
		}
	}

	logger := logrus.WithField(f.key, mysql_key)
	if strings.Contains(msg, "[info]") {
		logger.Info(data...)
	} else if strings.Contains(msg, "[warn]") {
		logger.Warn(data...)
	} else if strings.Contains(msg, "[error]") {
		logger.Error(data...)
	} else {
		logger.Warn(data...)
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
