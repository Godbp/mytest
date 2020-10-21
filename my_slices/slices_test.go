package my_url

import (
	"fmt"
	"testing"
)

func TestMySlices(t *testing.T) {
	s1, s2 := MySlices()
	fmt.Printf("s1=[%v]", s1)
	fmt.Printf("s2=[%v]", s2)
}
