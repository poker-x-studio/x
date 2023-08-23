/*
功能：goroutine
说明:
*/
package xgo_manager

import (
	"fmt"

	"github.com/poker-x-studio/x/xdone"
	"github.com/poker-x-studio/x/xservice"
)

type GoHandler func()

// xGo goroutine信息
type xGo struct {
	go_handler GoHandler         //handler
	goid       int64             //id
	status     xservice.Status   //状态
	done       xdone.DoneChannel //停止channel
}

// String 转化为字符串
func (g *xGo) String() string {
	return fmt.Sprintf("[xGo,goid:%d,status:%s,done:%+v]", g.goid, g.status.String(), g.done)
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
