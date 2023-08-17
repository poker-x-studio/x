/*
功能：计时器封装
说明：[计时器计数]默认一直增加，在调用处理函数后，返回的类型，可以重置 [计时器计数]
*/
package xticker

import (
	"fmt"
	"sync"
	"time"

	"github.com/poker-x-studio/x/xdebug"
	"github.com/poker-x-studio/x/xdone"
	"github.com/poker-x-studio/x/xlog"
	"github.com/poker-x-studio/x/xservice"
)

// 接口
type ITickerHandler interface {
	On_ticker(ticker_count int) TICKER_COUNT_TYPE
}

var ticker_index int //测试使用

type Ticker struct {
	xservice.Service                   //服务基类
	done             xdone.DoneChannel //停止channel
	index            int               //测试使用
	ticker           *time.Ticker      //
	ticker_count     int               //ticker计数
	duration         int               //时间间隔,秒为单位
	i_ticker_handler ITickerHandler    //接口对象
	wg               sync.WaitGroup    //
}

func NewTicker(i_ticker_handler ITickerHandler, duration int) *Ticker {
	ticker := &Ticker{}
	ticker.i_ticker_handler = i_ticker_handler
	ticker.duration = duration
	ticker_index++ //测试使用
	ticker.index = ticker_index
	return ticker
}

// 服务开始
func (t *Ticker) Service_start(ticker_count int) error {
	if t.i_ticker_handler == nil {
		return fmt.Errorf("计时器,接口设置错误")
	}
	if xdebug.Is_debug() {
		fmt.Printf("[Ticker],Service_start(),t.index:%d,goid:%d\r\n", t.index, xdebug.Go_id())
	}

	if t.Is_running() {
		return fmt.Errorf("计时器,服务已经开启")
	}

	t.ticker = time.NewTicker(time.Duration(t.duration) * time.Second)
	t.With_lock(func() {
		t.ticker_count = ticker_count
	})
	t.Update_status(xservice.STATUS_RUNNING)

	t.wg.Add(1)
	go t.ticker_handle()
	return nil
}

// ticker处理
func (t *Ticker) ticker_handle() {
	xlog_entry := xlog.New_entry(xdebug.Funcname())
	xlog_entry.Tracef("[Ticker],goid:%d", xdebug.Go_id())

	defer t.wg.Done()
	defer t.stop()

	for {
		select {
		case <-t.done.Done(): //停止，退出goroutine
			xlog_entry.Trace("goroutine,exit")
			return
		case <-t.ticker.C: //触发
			//xlog_entry.Tracef("ticker_count:%d", t.ticker_count)

			//接口调用
			if t.i_ticker_handler != nil {
				ticker_count_type := t.i_ticker_handler.On_ticker(t.ticker_count)
				if ticker_count_type == RESET {
					t.ticker_count = 0
				} else {
					t.ticker_count++
				}
			} else {
				t.ticker_count++
			}
		}
	}
}

// 服务停止
// 必须新开goroutine来关闭
// 如果不新开goroutine的话，调用方可能在 _ticker_handle 中的计时器处理函数
// 直接调用Service_stop()，导致 同一个goroutine 一个读 stop_ch，一个写 stop_ch
// 形成死锁
func (t *Ticker) Service_stop() error {
	go t.service_stop()
	return nil
}

// 服务停止
func (t *Ticker) service_stop() error {
	if t.Is_dead() {
		return fmt.Errorf("计时器,服务已经关闭")
	}
	t.Update_status(xservice.STATUS_DEAD)

	t.done.Notify_done()
	t.wg.Wait()    //等待退出goroutine
	t.done.Close() //sender负责关闭
	return nil
}

// 计时器停止(内部处理函数)
func (t *Ticker) stop() {
	fmt.Printf("[Ticker],stop(),t.index:%d,goid:%d\r\n", t.index, xdebug.Go_id())

	t.With_lock(func() {
		t.ticker.Stop()
		t.ticker = nil
		t.ticker_count = 0
	})
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
