/*
功能：
说明：
https://juejin.cn/post/6844903549881909256
logrus默认是没有像logrotate那样的日志切割功能
*/
package local_file_hook

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	"github.com/poker-x-studio/x"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

const (
	ROTATION_TIME  = time.Hour * 24
	ROTATION_COUNT = 30
)

var _DefaultFormatter = &logrus.TextFormatter{
	TimestampFormat: x.DATE_FORMAT, //时间格式
	FullTimestamp:   true,
	DisableColors:   true}

// LocalFileHook ...
type LocalFileHook struct {
	levels []logrus.Level
	writer *rotatelogs.RotateLogs
	pool   sync.Pool
}

// 本地异步循环日志
func NewLocalFileHook(dir, filename string, levels []logrus.Level) *LocalFileHook {
	os.Mkdir(dir, os.ModePerm)

	dir, err := filepath.Abs(dir)
	if err != nil {
		panic(err)
	}
	// Abs 会调用 Clean 方法, 因此会去除dir结尾的“/”
	dir += "/"

	options := []rotatelogs.Option{
		// WithRotationTime设置日志分割的时间,这里设置为一天分割一次
		rotatelogs.WithRotationTime(ROTATION_TIME),
		// WithMaxAge和WithRotationCount二者只能设置一个,
		// WithMaxAge设置文件清理前的最长保存时间,
		// WithRotationCount设置文件清理前最多保存的个数.
		//rotatelogs.WithMaxAge(time.Hour*24),
		rotatelogs.WithRotationCount(ROTATION_COUNT),
	}
	// WithLinkName为最新的日志建立软连接,以方便随着找到当前日志文件
	option := rotatelogs.WithLinkName(dir + filename + ".log")

	if runtime.GOOS == x.LINUX {
		options = append(options, option)
	}

	writer, err := rotatelogs.New(
		dir+filename+"_%Y_%m_%d.log",
		options...,
	)
	if err != nil {
		panic(fmt.Errorf("config local file system for logger error:%v", err))
	}

	hook := &LocalFileHook{
		levels: levels,
		writer: writer,
	}
	hook.pool.New = func() interface{} {
		return new(bytes.Buffer)
	}

	return hook
}

func (h *LocalFileHook) Fire(entry *logrus.Entry) error {
	buffer := h.pool.Get().(*bytes.Buffer)
	buffer.Reset()
	defer h.pool.Put(buffer)

	entry.Buffer = buffer
	serialized, err := _DefaultFormatter.Format(entry)
	if err != nil {
		return err
	}
	entry.Buffer = nil
	_, err = h.writer.Write(serialized)

	// fmt.Print(string(serialized))
	return err
}

func (h *LocalFileHook) Levels() []logrus.Level {
	return h.levels
}

func (h *LocalFileHook) Close() {
	h.writer.Close()
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
