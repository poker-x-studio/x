/*
功能：解析时间
说明：
*/
package xtime

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/poker-x-studio/x"
)

// 解析时间
// 格式要求 12:30 或 10:00 或 0:00 或 23:00
func Time_parse(t string) (*time.Time, error) {
	split_txts := strings.Split(t, ":")
	if len(split_txts) < 2 {
		return nil, errors.New("时间格式错误,正确举例:12:30 或 10:00 或 0:00 或 23:00")
	}

	//小时
	h, err := strconv.Atoi(split_txts[0])
	if err != nil {
		return nil, err
	}

	//分钟
	min, err := strconv.Atoi(split_txts[1])
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	y := now.Year()
	m := now.Month()
	d := now.Day()

	str_time := fmt.Sprintf("%d-%02d-%02d %02d:%02d:00", y, m, d, h, min)

	loc, err := time.LoadLocation("UTC")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	the_time, err := time.ParseInLocation(x.DATE_FORMAT, str_time, loc)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &the_time, nil
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
