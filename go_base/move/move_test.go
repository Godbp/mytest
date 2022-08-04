package move

import (
	"fmt"
	"testing"
)

var n = 10

const N = 10

func TestMove(t *testing.T) {
	var nn = 1 << n
	var NN = 1 << N
	var x = nn % 100
	var y = NN % 100
	fmt.Printf("x=%d \ny=%d", x, y)
}
