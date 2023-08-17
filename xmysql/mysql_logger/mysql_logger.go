/*
功能：mysql操作对应的文件日志
说明：
*/
package mysql_logger

import (
	"fmt"
	"path"
	"time"

	"github.com/poker-x-studio/x/xdebug"
	"github.com/poker-x-studio/x/xpath"
	"github.com/poker-x-studio/x/xutils"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
)

// 创建文件
func New_mysql_logger(filename string) (logger.Interface, func(), error) {
	log_level := logger.Warn //release日志等级3
	if xdebug.Is_debug() {   //debug日志等级4
		log_level = logger.Info
	}

	//目录
	executable_dir := xpath.Executable_dir()
	log_dir := path.Join(executable_dir, xutils.LOG_FOLDER)

	err := xpath.Mkdir(log_dir)
	if err != nil {
		panic("New_mysql_logger(),创建目录失败,log_dir:" + log_dir + ",err:" + err.Error())
	}
	logrus.Infof("New_mysql_logger(),创建目录,log_dir:%s", log_dir)

	file_logger, file_logger_close_handler, err := new_mysql_logger(log_dir, filename, log_level)
	if err != nil {
		panic("new mysql log failed, err:%s" + err.Error())
	}
	return file_logger, file_logger_close_handler, err
}

func new_mysql_logger(log_dir string, log_filename string, loglevel logger.LogLevel) (logger.Interface, func(), error) {
	if log_filename == "" {
		return nil, nil, fmt.Errorf("log_filename is nil")
	}

	file_logger, file_logger_close_handler, err := newFileLogger(log_dir, log_filename)
	if err != nil {
		return nil, nil, fmt.Errorf("new file log failed, err:%s", err.Error())
	}

	cfg := logger.Config{
		SlowThreshold: time.Second,
		LogLevel:      loglevel,
		Colorful:      true,
	}
	mylogger := logger.New(file_logger, cfg)

	return mylogger, file_logger_close_handler, nil
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
