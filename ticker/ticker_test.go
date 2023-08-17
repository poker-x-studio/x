/*
功能：测试单元
说明：
*/
package ticker

import (
	"fmt"
	"testing"
	"time"

	"github.com/poker-x-studio/x/xdebug"
)

type TickerHandler struct {
}

func (t *TickerHandler) On_ticker(ticker_count int) TICKER_COUNT_TYPE {
	fmt.Printf("On_ticker(),ticker_count:%d,goid:%d\r\n", ticker_count, xdebug.Go_id())

	return INCREASE
}

func Test(t *testing.T) {

	ticker_handler := &TickerHandler{}
	tickers := make([]*Ticker, 0)
	ticker_cnt := 5

	//多个计时器
	for i := 0; i < ticker_cnt; i++ {
		ticker := NewTicker(ticker_handler, 1)
		tickers = append(tickers, ticker)
	}

	for i := 0; i < ticker_cnt; i++ {
		tickers[i].Service_start(0)
	}

	time.Sleep(10 * time.Second)

	for i := 0; i < ticker_cnt; i++ {
		tickers[i].Service_stop()
	}

	time.Sleep(15 * time.Second)
	fmt.Println("end")
}

func Test1(t *testing.T) {
	var ch chan struct{}
	fmt.Println(ch)

	ch = make(chan struct{}, 0)
	fmt.Println(ch)

	ch = make(chan struct{}, 10)
	fmt.Println(ch)

}

//-----------------------------------------------
//					the end
//-----------------------------------------------
