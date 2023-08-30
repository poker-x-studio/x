/*
功能：
说明：
*/
package xutils

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

// Wait_for_signal 等待进程结束信号
func Wait_for_signal() {
	//进程结束,释放资源
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	signal := <-ch

	logrus.WithField("signal", signal).Info("server finish.")
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
