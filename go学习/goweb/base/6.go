package main

import "fmt"

func main() {
	/*
			*php 数组  go对应循环
			PHP
		     $data = ['a', 'b', 'c'];
		    foreach($data as $index => $value)
		    echo $index . $value . '|' ; // 輸出：0a|1b|2c|

		    foreach($data as $index => $value)
		    echo $index . '|' ; // 輸出：0|1|2|

		    foreach($data as $value)
		    echo $value . '|' ; // 輸出：a|b|c|
	*/

	data := []string{"a", "b", "c"}

	for index, value := range data {
		fmt.Printf("%d%s|", index, value) // 輸出：0a|1b|2c|
	}

	for index := range data {
		fmt.Printf("%d|", index) // 輸出：0|1|2|
	}

	for _, value := range data {
		fmt.Printf("%s|", value) // 輸出：a|b|c|
	}
}
