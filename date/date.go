package date

import (
	"time"
)

const (
	SendQWMsgDuration int64 = 10 * 60
	QWRemindStart     int64 = 9
	QWRemindEnd       int64 = 22
	OnyDay int64 = 24 * 60 * 60
)

// GetQWFriendRemind 判断是否需要企微提醒
func GetQWFriendRemind(qwRemindEnd, qwRemindStart, qwRemindDuration int64) (nextRemindTime int64, isInstantlySend bool) {
	if qwRemindEnd == 0 {
		qwRemindEnd = QWRemindEnd
	}
	if qwRemindStart == 0 {
		qwRemindStart = QWRemindStart
	}
	if qwRemindDuration == 0 {
		qwRemindDuration = SendQWMsgDuration
	}
	rs := GetTimer(qwRemindStart * 60 * 60)
	re := GetTimer(qwRemindEnd * 60 * 60)
	now := time.Now().Unix()
	if rs > now {
		return rs, isInstantlySend
	}
	if re < now {
		return rs + OnyDay, isInstantlySend
	}
	isInstantlySend = true
	return now + qwRemindDuration, isInstantlySend
}

// GetTimer 获取一天中的摸个时间点的时间戳
func GetTimer(ti int64) int64 {
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	return t.Unix() + ti
}