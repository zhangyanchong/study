package main

import (
	"fmt"
	"time"
)

/*
*并发
 */
func cal(a int, b int) {
	err := recover()
	if err != nil {
		fmt.Println("add ele fail")
	}
	c := a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
}

func main() {
	num := runtime.NumCPU()
	fmt.Printf(num)

	for i := 0; i < 10; i++ {
		go cal(i, i+1) //启动10个goroutine 来计算
	}
	time.Sleep(time.Second * 2) // sleep作用是为了等待所有任务完成  基本靠猜
}
