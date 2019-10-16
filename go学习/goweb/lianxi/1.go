package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

//Person定义了方法
func (p *Person) PrintInfo() {
	fmt.Printf("%s,%c,%d\n", p.name, p.sex, p.age)
}

type Student struct {
	Person // 匿名字段，那么Student包含了Person的所有字段
	id     int
	addr   string
}

func main() {
	p := Person{"mike", 'm', 18}
	p.PrintInfo()

	s := Student{Person{"yoyo", 'f', 20}, 2, "sz"}
	s.PrintInfo()
}
