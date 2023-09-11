/*
功能：labstack_log
说明：使用gommon包,也是echo框架使用的logger
*/
package labstack_log

import (
	"io"
	"os"
	"path"

	"github.com/labstack/gommon/log"
	"github.com/poker-x-studio/x/xpath"
	"github.com/poker-x-studio/x/xutils"
)

// 初始化日志
func Init_logger(debug bool, filename string) (xutils.HandlerClose, error) {
	dir := xpath.Executable_dir()

	//创建目录
	log_path := path.Join(dir, xutils.LOG_FOLDER)
	os.Mkdir(log_path, os.ModePerm)

	filepath := path.Join(dir, xutils.LOG_FOLDER, filename)
	file, err := os.Create(filepath)
	if err != nil {
		return nil, err
	}
	close := func() {
		file.Close()
	}

	//输出到文件和console
	writers := io.MultiWriter(os.Stdout, file)
	log.SetOutput(writers)

	//设置日志等级
	if debug {
		log.EnableColor()
		log.SetLevel(log.DEBUG)
	} else {
		log.SetLevel(log.INFO)
	}

	header := `time=${time_rfc3339_nano} level=${level} file=${long_file} line=${line}`
	log.SetHeader(header)
	return close, nil
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
