package async_hook

import (
	"errors"
	"sync"
	"sync/atomic"

	"github.com/sirupsen/logrus"
)

const (
	_AsyncQueueLen = 4096
)

// FilterFunc 返回 true 表示通过，false 表示被过滤。
type Filter func(entry *logrus.Entry) bool

// Firer 写入函数
type Firer func(entry *logrus.Entry) error

// Hook for logrus，异步写入 sql db，退出程序时调用 Close() 确保日志完全写入完毕。
type AsyncHook struct {
	levels        []logrus.Level
	originalFirer Firer
	filter        Filter

	logData chan *logrus.Entry
	closed  int32
	wg      sync.WaitGroup
}

// NewAsyncHook ...
func NewAsyncHook(levels []logrus.Level, firer Firer, filter Filter) *AsyncHook {
	h := &AsyncHook{
		levels:        levels,
		originalFirer: firer,
		filter:        filter,
		logData:       make(chan *logrus.Entry, _AsyncQueueLen),
	}

	go func() {
		for entry := range h.logData {
			h.originalFirer(entry)
			h.wg.Done()
		}
	}()

	return h
}

// NewAsyncHookWithHook 将现有的 logrus.Hook 转换成异步形式.
func NewAsyncHookWithHook(hook logrus.Hook, filter Filter) *AsyncHook {
	return NewAsyncHook(hook.Levels(), hook.Fire, filter)
}

// Levels logrus.Hook interface
func (h *AsyncHook) Levels() []logrus.Level {
	return h.levels
}

// Fire logrus.Hook interface
func (h *AsyncHook) Fire(entry *logrus.Entry) error {
	if atomic.LoadInt32(&h.closed) != 0 {
		return errors.New("asynchook closed")
	}

	if h.filter != nil && !h.filter(entry) {
		return nil
	}

	newe := logrus.NewEntry(entry.Logger)
	for k, v := range entry.Data {
		newe.Data[k] = v
	}
	newe.Time = entry.Time
	newe.Caller = entry.Caller
	newe.Level = entry.Level
	newe.Message = entry.Message

	h.wg.Add(1)
	h.logData <- newe

	return nil
}

// Close 关闭异步记录循环, 该调用会等待所有操作完成.
func (h *AsyncHook) Close() {
	if atomic.LoadInt32(&h.closed) == 0 {
		atomic.StoreInt32(&h.closed, 1)
	}

	close(h.logData)

	h.wg.Wait()
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
