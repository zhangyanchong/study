package main

import "fmt"

func main() {
	/*
	*数组和切片
	 */
	var shuzu = [4]int{1, 2, 3, 4}
	fmt.Println(shuzu[0])

	var zifu = [2]string{"haha", "mao"}
	fmt.Println(zifu[1])

	/*
	*切片
	 */
	fmt.Println(shuzu[2:3])
}
