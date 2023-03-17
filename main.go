package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("====== 开始测试死循环 ==========")

	for i := 0; ; i++ {
		fmt.Println(fmt.Sprintf("========= %d ==========\n", i))

		time.Sleep(time.Second * 1)
	}
}
