package myTime

import (
	"fmt"
	"testing"
)

func TestTime1(t *testing.T) {
	fmt.Printf("%d\n", GetNowTimeDay("8:30"))
	tt := []string{"8:30", "9:30", "12:30", "11:05", "3:05", "16:01", "16:00"}
	fmt.Printf("%d", GetNearNoticeTime(tt))
}
