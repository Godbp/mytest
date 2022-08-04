package channel

import (
	"fmt"
	"time"
)

// Consumer ...
func Consumer() {
	var num = 0
	for {
		num += 1
		time.Sleep(1 * time.Second)
		fmt.Printf("开始写入 %d \n", num)
		ch <- num
		if num == 100 {
			close(ch)
			return
		}
	}
}
