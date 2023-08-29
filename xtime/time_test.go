/*
功能：测试单元
说明：
*/
package xtime

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func Test_1(t *testing.T) {
	tm := "12:30"
	tm = "23:00"
	//tm = "0:00"
	time1, err := Time_parse(tm)
	if err != nil {
		return
	}
	fmt.Println(time1.Unix())
	/*
	   timestamp := time.Now().Unix()
	   Timestamp_2_local_time(timestamp)
	   Timestamp_2_utc_time(timestamp)
	*/
}

func Test_2(t *testing.T) {

	// Add 时间相加
	now := time.Now()
	fmt.Println("now:", now.String())

	// ParseDuration parses a duration string.
	// A duration string is a possibly signed sequence of decimal numbers,
	// each with optional fraction and a unit suffix,
	// such as "300ms", "-1.5h" or "2h45m".
	//  Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	// 10分钟前
	m, _ := time.ParseDuration("-1m")
	m1 := now.Add(m)
	fmt.Println(m1)
	fmt.Println("now:", now.String())

	// 8个小时前
	h, _ := time.ParseDuration("-1h")
	h1 := now.Add(8 * h)
	fmt.Println(h1)
	fmt.Println("now:", now.String())

	// 一天前
	d, _ := time.ParseDuration("-24h")
	d1 := now.Add(d)
	fmt.Println(d1)

	printSplit(50)

	// 10分钟后
	mm, _ := time.ParseDuration("1m")
	mm1 := now.Add(mm)
	fmt.Println(mm1)

	// 8小时后
	hh, _ := time.ParseDuration("1h")
	hh1 := now.Add(hh)
	fmt.Println(hh1)

	// 一天后
	dd, _ := time.ParseDuration("24h")
	dd1 := now.Add(dd)
	fmt.Println(dd1)

	printSplit(50)

	// Sub 计算两个时间差
	subM := now.Sub(m1)
	fmt.Println(subM.Minutes(), "分钟")

	sumH := now.Sub(h1)
	fmt.Println(sumH.Hours(), "小时")

	sumD := now.Sub(d1)
	fmt.Printf("%v 天\n", sumD.Hours()/24)

}

func printSplit(count int) {
	fmt.Println(strings.Repeat("#", count))
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
