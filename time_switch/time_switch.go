/*
功能：时间开关
说明：

1 时间字符串规则 10:22,20:00
2 stop_time<start_time,结束时间是下一天的时间点
*/

package time_switch

import (
	"fmt"
	"time"
	"x/xdebug"
	"x/xlog"
	"x/xservice"
	"x/xtime"
)

// 时间开关
type TimeSwitch struct {
	xservice.Service              //
	stop_time        time.Time    //停止时间,24小时制,比如 3:45,UTC时间,
	start_time       time.Time    //开始时间,24小时制,比如 4:45,UTC时间,必须满足stop_time<start_time
	handler_stop     HandlerStop  //停止函数
	handler_start    HandlerStart //开始函数
	save_timer       *time.Timer  //保存定时器
}

// 创建
func NewTimeSwitch(stop_time string, start_time string, options ...Option) *TimeSwitch {
	t := &TimeSwitch{}
	for _, option := range options {
		option(t)
	}

	if err := t.init(stop_time, start_time); err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return t
}

// 初始化
func (t *TimeSwitch) init(stop_time string, start_time string) error {
	xlog_entry := xlog.New_entry(xdebug.Funcname())

	//解析时间
	stop, err := xtime.Time_parse(stop_time)
	if err != nil {
		return fmt.Errorf("错误,时间配置,stop_time:%s,err:%s,", stop_time, err.Error())
	}
	xlog_entry.Tracef("stop:%s", stop.String())

	start, err := xtime.Time_parse(start_time)
	if err != nil {
		return fmt.Errorf("错误,时间配置,start_time:%s,err:%s,", start_time, err.Error())
	}
	xlog_entry.Tracef("start:%s", start.String())

	//开始时间需要大于结束时间
	if start.Before(*stop) {
		return fmt.Errorf("错误,时间配置,stop必须小于start,start_time:%s,stop_time:%s,", start_time, stop_time)
	}

	t.stop_time = *stop
	t.start_time = *start
	return nil
}

// 服务启动
func (t *TimeSwitch) Service_start() {
	if t.Is_running() {
		return
	}
	t.kill_timer()
	//准备自动启动
	t.ready_auto_start()
}

// 服务停止
func (t *TimeSwitch) Service_stop() {
	if t.Is_dead() {
		return
	}
	t.kill_timer()
	t.auto_stop()
}

// 准备自动启动
func (t *TimeSwitch) ready_auto_start() {
	xlog_entry := xlog.New_entry(xdebug.Funcname())
	xlog_entry.Info("[TimeSwitch]ready_auto_start,准备自动启动")

	time_now := time.Now().UTC()
	xlog_entry.Infof("time_now:%s", time_now.String())

	//当前时间是否在停止时间段
	is_stop_time := false
	if time_now.After(t.stop_time) && time_now.Before(t.start_time) {
		is_stop_time = true
	}

	if is_stop_time { //停止时间段
		diff_time := t.start_time.Sub(time_now)
		seconds := int(diff_time.Seconds())
		h, m := xtime.Seconds_2_hh_mm(seconds)
		xlog_entry.Infof("[TimeSwitch]停止时间段,还没到开始时间,%d小时%d分钟后启动", h, m)
		t.save_timer = time.AfterFunc(time.Duration(seconds)*time.Second, func() {
			t.kill_timer()
			t.auto_start()
		})
	} else { //非停止时间段
		xlog_entry.Info("[TimeSwitch]非停止时间段,立即启动")
		go t.auto_start()
		return
	}
}

// 自动启动
func (t *TimeSwitch) auto_start() {
	if t.Is_running() {
		return
	}

	xlog_entry := xlog.New_entry(xdebug.Funcname())
	xlog_entry.Info("[TimeSwitch]auto_start,自动启动")

	if t.handler_start != nil {
		t.handler_start()
		t.Update_status(xservice.STATUS_RUNNING)
	}

	time_now := time.Now().UTC()
	xlog_entry.Infof("time_now:%s", time_now.String())

	//计算下次停止时间
	var new_time_stop time.Time
	if time_now.After(t.stop_time) { //关闭时间已经过去，则第二天关闭
		new_time_stop = t.stop_time.AddDate(0, 0, 1)
	} else { //关闭时间还没过去
		new_time_stop = t.stop_time
	}
	xlog_entry.Infof("new_time_stop:%s", new_time_stop.String())

	diff_time := new_time_stop.Sub(time_now)
	seconds := int(diff_time.Seconds())
	h, m := xtime.Seconds_2_hh_mm(seconds)
	xlog_entry.Infof("%d小时%d分钟后停止", h, m)

	t.save_timer = time.AfterFunc(time.Duration(seconds)*time.Second, func() {
		t.kill_timer()
		t.auto_stop()
	})
}

// 自动停止
func (t *TimeSwitch) auto_stop() {
	if t.Is_dead() {
		return
	}

	xlog_entry := xlog.New_entry(xdebug.Funcname())
	xlog_entry.Info("[TimeSwitch]auto_stop,自动停止")

	if t.handler_stop != nil {
		t.handler_stop()
		t.Update_status(xservice.STATUS_DEAD)
	}

	time_now := time.Now().UTC()
	xlog_entry.Infof("time_now:%s", time_now.String())

	//计算下次开始时间
	diff_time := t.start_time.Sub(time_now)
	seconds := int(diff_time.Seconds())
	h, m := xtime.Seconds_2_hh_mm(seconds)
	xlog_entry.Infof("%d小时%d分钟后启动", h, m)

	t.save_timer = time.AfterFunc(time.Duration(seconds)*time.Second, func() {
		t.kill_timer()
		t.auto_start()
	})
}

// 人工启动
func (t *TimeSwitch) Manual_start() {
	if t.Is_running() {
		return
	}

	xlog_entry := xlog.New_entry(xdebug.Funcname())
	xlog_entry.Info("[TimeSwitch]Manual_start,人工启动")

	t.kill_timer()

	if t.handler_start != nil {
		t.handler_start()
		t.Update_status(xservice.STATUS_RUNNING)
	}
}

// 人工停止
func (t *TimeSwitch) Manual_stop() {
	if t.Is_dead() {
		return
	}
	xlog_entry := xlog.New_entry(xdebug.Funcname())
	xlog_entry.Info("[TimeSwitch]Manual_stop,人工停止")

	t.kill_timer()

	if t.handler_stop != nil {
		t.handler_stop()
		t.Update_status(xservice.STATUS_DEAD)
	}
}

// 销毁定时器
func (t *TimeSwitch) kill_timer() {
	if t.save_timer != nil {
		t.save_timer.Stop()
	}
	t.save_timer = nil
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
