package myTime

import (
	"fmt"
	"git.dustess.com/mk-base/util/date"
	"sort"
	"time"
)

// GetDateFormat 根据时间类型转换格式化日期
func GetDateFormat(t time.Time) string {
	return t.Format("2006-01-02")
}

// GetNowTimestampByString 通过string转换成时间戳
func GetNowTimestampByString(dataStr string, format string) int64 {
	t, _ := time.ParseInLocation(format, dataStr, time.Local)
	return t.Unix()
}

// GetNowTimeDay 获取传入时间在当前天的时间戳
func GetNowTimeDay(hm string) int64 {
	now := time.Now().AddDate(0, 0, 1)
	ns := date.GetDateFormat(now)
	ns = fmt.Sprintf("%s %s", ns, hm)
	t, _ := time.ParseInLocation("2006-01-02 15:04", ns, time.Local)
	return t.UnixMilli()
}

type IntSlice []int64

func (s IntSlice) Len() int { return len(s) }

func (s IntSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }

// GetNearNoticeTime 获取最近的通知时间 times=[hh:mm, hh:mm]
func GetNearNoticeTime(times []string) int64 {
	now := date.GetNowTimestampms()

	if len(times) == 0 {
		return 0
	}

	ti := make([]int64, 0, len(times))
	for _, t := range times {
		ti = append(ti, GetTimeDay(t, 0, 0, 0))
	}
	sort.Sort(IntSlice(ti))

	// 获取最后一个或者唯一一个数据与当前时间对比 如果小于当前时间，即说明今日提醒时间已过，需要设置第二天的第一个时间
	if ti[len(times)-1] < now {
		return ti[0] + 24*60*60*1000
	}

	for _, tt := range ti {
		if now > tt {
			continue
		}
		return tt
	}
	return 0
}

// GetTimeDay 获取传入时间在当前天的时间戳 y 年份偏移量 m 月份偏移量 d 日期偏移量
func GetTimeDay(hm string, y, m, d int) int64 {
	now := time.Now().AddDate(y, m, d)
	ns := date.GetDateFormat(now)

	ns = fmt.Sprintf("%s %s", ns, hm)
	t, _ := time.ParseInLocation("2006-01-02 15:04", ns, time.Local)
	return t.UnixMilli()
}
