package my_slice

import (
	"fmt"
	"strconv"
	"testing"
)

func TestMySlices(t *testing.T) {
	//s1, s2 := MySlices2()
	//fmt.Printf("s1=[%v]", s1)
	//fmt.Printf("s2=[%v]", s2)
	//l1 := []string{"etnBvEDAAABHV86PCNAFSYnas3XhZ0Zw", "etnBvEDAAAQ5ysLNA_h0TvETbXyRazvw", "etnBvEDAAA7aKmCRJ6P57QfKg6u-lQNw", "etnBvEDAAAV7TaCiiL_PIh64vpedFZag"}
	//l2 := []string{"etnBvEDAAA7aKmCRJ6P57QfKg6u-lQNw", "etnBvEDAAAV7TaCiiL_PIh64vpedFZag"}
	//r1, r2 := MatchingSlices(l1, l2)
	//fmt.Printf("%+v \n%+v", r1, r2)
}

func TestAppendAddress(t *testing.T) {
	ls := make([]AppendAddressStrut, 0, 8)
	for i := 0; i < 9; i++ {
		ls = append(ls, AppendAddressStrut{Name: strconv.Itoa(i)})
	}
	res := AppendAddress(ls)
	fmt.Printf("%+v", res)
}
