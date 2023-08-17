/*
功能：测试单元
说明：
*/
package time_switch

import (
	"fmt"
	"testing"
	"time"
)

func TestXxx(t *testing.T) {
	options := []Option{
		WithStopHandler(stop),
		WithStarHandler(start),
	}

	tc := NewTimeSwitch("2:45", "3:45", options...)
	if tc == nil {
		return
	}
	tc.Service_start()

	time.Sleep(5 * time.Second)
	tc.Service_start()

	time.Sleep(5 * time.Second)
	tc.Service_stop()

	time.Sleep(5 * time.Second)
	tc.Service_start()

	time.Sleep(5 * time.Second)
	tc.Service_stop()
}

func stop() {
	fmt.Println("stop()函数")
}

func start() {
	fmt.Println("start()函数")
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
