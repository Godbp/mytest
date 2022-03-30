package myString

import "fmt"

// Slice 支持负数转换
func Slice(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

// SliceList 数组切片
func SliceList(l []string, slice int) []string {
	length := len(l) / slice
	y := len(l) % slice
	if y > 0 {
		length += 1
	}
	for i := 0; i < length; i++ {
		star := i * slice
		end := (i + 1) * slice
		if end > len(l) {
			end = len(l)
		}
		if star > len(l) {
			return nil
		}
		fmt.Printf("%d === %+v", i, l[star:end])
	}
	return nil
}
