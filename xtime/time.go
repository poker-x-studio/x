/*
功能：时间
说明：
*/
package xtime

import (
	"time"
)

// Seconds_2_hh_mm 秒转小时分钟
func Seconds_2_hh_mm(seconds int) (int, int) {
	h := seconds / 3600
	m := (seconds - h*3600) / 60
	return h, m
}

/*
// 时间戳转local时间

	func Timestamp_2_local_time(timestamp int64) time.Time {
		//时间戳转time
		tm := time.Unix(timestamp, 0)
		tm_str := tm.Format(LAYOUT)
		fmt.Println(tm_str)
		return tm
	}
*/

// Timestamp_2_utc_time 时间戳转utc时间
func Timestamp_2_utc_time(timestamp int64) time.Time {
	//const func_tag = "Timestamp_2_utc_time(),"

	//时间戳转time
	tm := time.Unix(timestamp, 0)
	//tm_str := tm.Format(x.DATE_FORMAT)
	//fmt.Println(func_tag, tm_str)

	//本地时间转utc时间
	utc_tm := tm.UTC()
	//utc_tm_str := utc_tm.Format(x.DATE_FORMAT)
	//fmt.Println(func_tag, utc_tm_str)

	return utc_tm
}

/*
// local当前时间
func Local_time_now() time.Time {
	timestamp := time.Now().Unix()
	local_tm := Timestamp_2_local_time(timestamp)
	return local_tm
}

// utc当前时间
func Utc_time_now() time.Time {
	timestamp := time.Now().Unix()
	utc_tm := Timestamp_2_utc_time(timestamp)
	return utc_tm
}
*/
//-----------------------------------------------
//					the end
//-----------------------------------------------
