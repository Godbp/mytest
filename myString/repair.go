package myString

import (
	"strconv"
	"time"
)

// TransDelayTime 转换delay时间
func TransDelayTime(delayTime string) (time.Duration, error) {
	res := 0 * time.Second
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
	d, err := strconv.Atoi(string([]byte(delayTime)[:3]))
	if err != nil {
		return res, err
	}
	h, err := strconv.Atoi(string([]byte(delayTime)[3:6]))
	if err != nil {
		return res, err
	}
	m, err := strconv.Atoi(string([]byte(delayTime)[6:]))
	if err != nil {
		return res, err
	}
	return time.Duration(d)*24*time.Hour + time.Duration(h)*time.Hour + time.Duration(m)*time.Minute, nil
}
