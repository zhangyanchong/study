package main

import "fmt"

func main() {
	/*
	     结构体
		**/
	type stu struct {
		id   int8
		name string
		sex  int
		age  string
	}
	var xuesheng stu
	xuesheng.id = 1
	xuesheng.name = "zhangsan"
	xuesheng.sex = 2
	xuesheng.age = "14"
	fmt.Println(xuesheng.name)
}
