package my_slice

import "fmt"

func MySlices() ([]string, []string) {
	var s1 []string
	s2 := make([]string, 0)
	return s1, s2
}

func MySlices2(sl []string) {
	limit := len(sl) / 2

	cs := make([]string, 0, len(sl))
	for i := 0; i < limit+1; i++ {
		if len(sl) < (i+1)*2 {
			cs = append(cs, sl[i*2:]...)
			continue
		}
		cs = append(cs, sl[i*2:(i+1)*2]...)
	}
	fmt.Printf("%+v", cs)
}

// MatchingSlices 判断新旧两个数组元素的差异
func MatchingSlices(old, new []string) ([]string, []string) {
	move := make([]string, 0, 0)
	add := make([]string, 0, 0)
	// 旧元素不在新数组-获取新数组删除的元素
	for _, t := range old {
		if !StringSliceIn(t, new) {
			move = append(move, t)
		}
	}
	// 新元素不在旧数组-获取新数组添加的元素
	for _, t := range new {
		if !StringSliceIn(t, old) {
			add = append(add, t)
		}
	}
	return move, add
}

// StringSliceIn 判断传入字符串是否存在数组里面
func StringSliceIn(m string, s []string) bool {
	for i := range s {
		if m == s[i] {
			return true
		}
	}
	return false
}

type AppendAddressStrut struct {
	Name string `json:"name"`
}

// AppendAddress 追加地址
func AppendAddress(ls []AppendAddressStrut) []interface{} {
	res := make([]interface{}, 0, len(ls))
	slSize := 2
	sl := make([]interface{}, 0, slSize)
	for index, i := range ls {
		i.Name = i.Name + "a"
		sl = append(sl, i)
		if len(sl) == slSize {
			res = append(res, sl)
			sl = []interface{}{}
			continue
		}
		if len(ls) == index+1 {
			res = append(res, sl)
		}
	}
	return res
}
