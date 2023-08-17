/*
功能：计时器计数类型
说明：
*/
package ticker

type TICKER_COUNT_TYPE int

const (
	ERROR    TICKER_COUNT_TYPE = 0 //错误
	RESET    TICKER_COUNT_TYPE = 1 //重置归0
	INCREASE TICKER_COUNT_TYPE = 2 //增加
)

//-----------------------------------------------
//					the end
//-----------------------------------------------
