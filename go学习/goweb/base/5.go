package main

import "fmt"

func main() {
	/*
	*map 注意类型
	 */
	data := make(map[string]string)
	data["username"] = "YamiOdymel"
	data["password"] = "2016 Spring"
	fmt.Println(data["username"]) // 輸出：YamiOdymel

	/*
	*接口 可以接受任何类型的值
	 */
	mixedData2 := make(map[string]interface{})
	mixedData2["username"] = "接口"
	mixedData2["time"] = 123456
	fmt.Println(mixedData2["username"])
}
