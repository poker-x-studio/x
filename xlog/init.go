/*
功能：日志相关
说明：
*/
package xlog

import (
	"fmt"
	"path"
	"strconv"

	"github.com/poker-x-studio/x/xdebug"
	"github.com/poker-x-studio/x/xpath"
	"github.com/poker-x-studio/x/xutils"

	"github.com/sirupsen/logrus"
)

const (
	TAG = "TAG"
)

// Set_level 设置日志等级
func set_level() {
	logrus.SetLevel(get_level())
}

// get_level 获取等级
func get_level() logrus.Level {
	//所有debug标志统一在xdebug中设置
	if xdebug.Is_debug() {
		return logrus.TraceLevel
	}
	return logrus.InfoLevel
}

// Init_logrus 初始化logrus
func Init_logrus(filename string) {
	set_level()

	//logrus.SetReportCaller(true) //输出文件，调用函数，行数，
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05", //时间格式
		FullTimestamp:   true,
	})

	executable_dir := xpath.Executable_dir()
	log_dir := path.Join(executable_dir, xutils.LOG_FOLDER)

	//创建文件夹
	xpath.Mkdir(log_dir)

	log_level := logrus.GetLevel()
	AddLocalFileHook(log_dir, filename, "", nil, log_level)

	logrus.Info("init")
}

// New_entry 新的入口
func New_entry(format string, a ...string) *logrus.Entry {
	tag := format
	if len(a) > 0 {
		tag = fmt.Sprintf(format, a)
	}
	return logrus.WithFields(logrus.Fields{
		TAG: tag,
	})
}

// New_entry_tag 新的入口
func New_entry_tag(tags ...string) *logrus.Entry {
	len := len(tags)
	tag_map := make(map[string]interface{}, 0)

	if len == 0 {
		tag_map[TAG] = TAG
	} else if len == 1 {
		tag_map[TAG] = tags
	} else {
		for i := 0; i < len; i++ {
			key := TAG + strconv.Itoa(i)
			tag_map[key] = tags[i]
		}
	}

	return logrus.WithFields(logrus.Fields(tag_map))
}

// New_entry_chat_id 新的入口
func New_entry_chat_id(tag string, chat_id int64) *logrus.Entry {
	return logrus.WithFields(logrus.Fields{
		TAG:       tag,
		"chat_id": chat_id,
	})
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
