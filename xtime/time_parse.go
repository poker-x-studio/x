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

	"github.com/poker-x-studio/x/xutils"
)

// Time_parse 解析时间
// 格式要求 12:30 或 10:00 或 0:00 或 23:00
func Time_parse(t string) (*time.Time, error) {
	txts := strings.Split(t, ":")
	if len(txts) != 2 {
		return nil, errors.New("时间格式错误,正确举例:12:30 或 10:00 或 0:00 或 23:00")
	}

	hm := make([]int, 0)
	for i := 0; i < len(txts); i++ {
		n, err := strconv.Atoi(txts[i])
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		hm = append(hm, n)
	}

	loc, err := time.LoadLocation("UTC")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	now := time.Now().UTC()
	str_time := fmt.Sprintf("%d-%02d-%02d %02d:%02d:00", now.Year(), now.Month(), now.Day(), hm[0], hm[1])
	local_time, err := time.ParseInLocation(xutils.DATE_FORMAT, str_time, loc)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &local_time, nil
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
