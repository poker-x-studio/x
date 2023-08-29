/*
功能：完成通道
说明：
*/
package xdone

// DoneChannel 完成通道
type DoneChannel struct {
	ch chan struct{}
}

// init 初始化
func (d *DoneChannel) init() {
	if d.ch == nil {
		d.ch = make(chan struct{}, 1)
	}
}

// Done 返回完成chan
func (d *DoneChannel) Done() <-chan struct{} {
	d.init()
	return d.ch
}

// Notify_done 通知结束
func (d *DoneChannel) Notify_done() {
	d.init()
	go func() {
		d.ch <- struct{}{}
	}()
}

// Close 关闭
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
