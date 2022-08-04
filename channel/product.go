package channel

import (
	"fmt"
)

func init() {
	go Product(ch)
}

func Product(ch chan int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("出错啦！！！")
			go Product(ch)
		}
	}()
	for {
		fmt.Printf("开始消费")
		d, ok := <-ch
		if !ok {
			return
		}
		if d/10 == 0 {
			panic("这是一个错误")
		}
		fmt.Printf("%s \n", d)
	}
}
