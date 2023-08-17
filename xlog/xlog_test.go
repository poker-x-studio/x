/*
功能：测试单元
说明：
*/
package xlog

import (
	"testing"
	"time"

	"github.com/poker-x-studio/x/xutils"

	"github.com/sirupsen/logrus"
)

func TestLog(t *testing.T) {
	//dir := "/home/danny/Documents/test/"
	//file := "filename"

	//windows
	dir := "C:\\_test_proj\\go\\Telegram\\cmd\\telegram_bot\\"
	file := "log"

	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: xutils.DATE_FORMAT, //时间格式
		FullTimestamp:   true,
	})

	logrus.SetReportCaller(true)

	AddLocalFileHook(dir, file, "", nil, logrus.TraceLevel)

	logrus.Info("========info")
	logrus.Warn("========warn")
	logrus.Error("========error")

	requestLogger := logrus.WithFields(logrus.Fields{"request_id": "request_id", "sadfsa": "dafsaf"})
	requestLogger.Info("something happened on that request")
	requestLogger.Warn("something not great happened")

	time.Sleep(time.Second)
}

func TestDDD(t *testing.T) {
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
