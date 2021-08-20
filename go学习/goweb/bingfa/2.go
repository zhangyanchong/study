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
	"runtime"
	"sync"
)

var  wg  sync.WaitGroup   //计数器

func hello(i int){
	fmt.Println("hello 娜扎",i)
	wg.Done()   //通知wg把计数器-1
}

func main(){ //开启一个主 goroutine 去执行main 函数
	runtime.GOMAXPROCS(4)  //使用几个cpu 核心去处理
	for i:=0;i<1000;i++{
		wg.Add(1)  //计数器加+1
		go hello(i)   //开启了一个goroutine 去执行hello这个函数
	}
	fmt.Println("hello main")
	wg.Wait()  //阻塞 等所有小弟都干完活之后才能收兵
}
