package my_slice

import "fmt"

func MySlices() ([]string, []string) {
	var s1 []string
	s2 := make([]string, 0)
	return s1, s2
}

func MySlices2(sl []string)  {
	limit := len(sl) / 2

	cs := make([]string, 0, len(sl))
	for i := 0; i < limit+1; i++ {
		if len(sl) < (i+1)*2{
			cs = append(cs, sl[i*2:]...)
			continue
		}
		cs = append(cs, sl[i*2:(i+1)*2]...)
	}
	fmt.Printf("%+v", cs)
}