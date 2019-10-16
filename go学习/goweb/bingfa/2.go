package main

/*
*使用sync包同步goroutine
sync大致实现方式
WaitGroup 等待一组goroutinue执行完毕. 主程序调用 Add 添加等待的goroutinue数量.
每个goroutinue在执行结束时调用 Done ，此时等待队列数量减1.，
主程序通过Wait阻塞，直到等待队列为0.
*/
import (
	"fmt"
	"sync"
)

func cal(a int, b int, n *sync.WaitGroup) {
	c := a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
	defer n.Done() //goroutinue完成后, WaitGroup的计数-1

}

func main() {
	var go_sync sync.WaitGroup //声明一个WaitGroup变量
	for i := 0; i < 10; i++ {
		go_sync.Add(1) // WaitGroup的计数加1
		go cal(i, i+1, &go_sync)
	}
	go_sync.Wait() //等待所有goroutine执行完毕
}
