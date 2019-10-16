package main

import "fmt"

/*
* 渠道学习
 */
func test(c chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("send ", i)
		c <- i
	}
}
func main() {
	ch := make(chan int)
	go test(ch)
	for j := 0; j < 10; j++ {
		fmt.Println("get ", <-ch)
	}
}
