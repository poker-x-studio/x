/*
功能：goroutine 管理
说明:
*/
package xgo_manager

import (
	"fmt"
	"sync"

	"github.com/poker-x-studio/x/xdebug"
	"github.com/poker-x-studio/x/xservice"
)

type xGoManager struct {
	xgo_map map[int64]*xGo //key为goid
	wg      sync.WaitGroup //
	mutex   sync.Mutex
}

var once sync.Once
var instance *xGoManager

// 单例
func Instance() *xGoManager {
	once.Do(func() {
		if instance == nil {
			instance = &xGoManager{}
			instance.init()
		}
	})
	return instance
}

// 初始化
func (mgr *xGoManager) init() {
	mgr.xgo_map = make(map[int64]*xGo, 0)
}

// 运行goroutine
func (mgr *xGoManager) Go(f func()) {
	mgr.mutex.Lock()
	defer mgr.mutex.Unlock()

	var xgo xGo
	var goid int64
	ch := make(chan int64)

	defer func() {
		xgo.goid = goid
		xgo.status = xservice.STATUS_RUNNING
		//fmt.Println(xgo.String())

		mgr.xgo_map[xgo.goid] = &xgo
		mgr.print()
		mgr.wg.Add(1)
	}()

	xgo.go_handler = func() {
		ch <- xdebug.Go_id() //获取goid
		f()
	}

	go xgo.go_handler()
	goid = <-ch
	close(ch)
}

// 结束了一个goroutine
func (mgr *xGoManager) Dead(goid int64) {
	mgr.mutex.Lock()
	defer mgr.mutex.Unlock()

	xgo, ok := mgr.xgo_map[goid]
	if !ok {
		return
	}
	if xgo.status == xservice.STATUS_DEAD {
		return
	}
	xgo.status = xservice.STATUS_DEAD
	//fmt.Println(xgo.String())
	mgr.print()
	mgr.wg.Done()
}

// 通知结束所有goroutine
func (mgr *xGoManager) Notify_all_dead() {
	mgr.mutex.Lock()
	defer mgr.mutex.Unlock()

	for _, v := range mgr.xgo_map {
		v.done.Notify_done()
	}

	go func() {
		mgr.wg.Wait()
		for _, v := range mgr.xgo_map {
			v.done.Close()
		}
	}()
}

// 结束的通道
func (mgr *xGoManager) Done(goid int64) <-chan struct{} {
	mgr.mutex.Lock()
	defer mgr.mutex.Unlock()

	xgo, ok := mgr.xgo_map[goid]
	if !ok {
		panic("")
	}
	//fmt.Println(xgo.String())
	return xgo.done.Done()
}

// 调试输出
func (mgr *xGoManager) print() {
	for _, v := range mgr.xgo_map {
		fmt.Println(v.String())
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
