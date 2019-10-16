package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://baijiahao.baidu.com/s?id=1638632403585158483"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	html, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		panic(err2)
	}
	// ReadAll返回[]byte
	fmt.Println(string(html))
}
