package main

import (
	"github.com/gogf/gf/frame/g"
	_ "xin/router"  //有用 加载路有用的
)

func main() {
	s := g.Server()
	s.Run()
}
