package myString

import (
	"strconv"
)

// 常量定义
const (
	Millisecond int64 = 1
	Second            = 1000 * Millisecond
	Minute            = 60 * Second
	Hour              = 60 * Minute
	Day               = 24 * Hour
)

// TransTimeFormat2Int 时间转换
func TransTimeFormat2Int(delayTime string) (int64, error) {
	repairStr := "000000000"
	strLen := 9
	// 向前补充缺失位数 "0"
	if len(delayTime) < strLen {
		delayTime = string([]byte(repairStr)[:strLen-len(delayTime)]) + delayTime
	}
	// 截取后9位
	if len(delayTime) > strLen {
		delayTime = string([]byte(repairStr)[strLen:])
	}
	d, err := strconv.ParseInt(string([]byte(delayTime)[:3]), 10, 64)
	if err != nil {
		return 0, err
	}
	h, err := strconv.ParseInt(string([]byte(delayTime)[3:6]), 10, 64)
	if err != nil {
		return 0, err
	}
	m, err := strconv.ParseInt(string([]byte(delayTime)[6:]), 10, 64)
	if err != nil {
		return 0, err
	}
	return d*Day + h*Hour + m*Minute, nil
}
