/*
功能：完成通道
说明：
*/
package xdone

type DoneChannel struct {
	ch chan struct{}
}

// 初始化
func (d *DoneChannel) init() {
	if d.ch == nil {
		d.ch = make(chan struct{}, 1)
	}
}

// 完成
func (d *DoneChannel) Done() <-chan struct{} {
	d.init()
	return d.ch
}

// 通知结束
func (d *DoneChannel) Notify_done() {
	d.init()
	go func() {
		d.ch <- struct{}{}
	}()
}

// 关闭
func (d *DoneChannel) Close() {
	if d.ch == nil {
		return
	}
	close(d.ch)
	d.ch = nil
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
