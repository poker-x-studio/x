/*
功能：生命周期
说明：
1 输出函数运行时间[单位秒]
*/
package xdebug

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

// LifeTime 生命周期
type LifeTime struct {
	start_timestamp int64 //开始时间戳
	end_timestamp   int64 //结束时间戳
}

// Start 开始
func (l *LifeTime) Start() {
	l.start_timestamp = time.Now().Unix()
	logrus.Tracef("start_timestamp:%d", l.start_timestamp)
}

// End 结束
func (l *LifeTime) End() {
	l.end_timestamp = time.Now().Unix()
	logrus.Tracef("end_timestamp:%d,%s", l.end_timestamp, l.string())
}

// 转字符串
func (l *LifeTime) string() string {
	diff := l.end_timestamp - l.start_timestamp

	min := diff / 60
	sec := diff - min*60
	str := fmt.Sprintf("生命周期,时间戳差值diff:%d(秒),时间=%d:%02d(%d分%d秒)", diff, min, sec, min, sec)
	return str
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
