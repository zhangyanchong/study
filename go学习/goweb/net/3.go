// 在本程序同目录下创建images文件夹，放入一些文件
// 打开浏览器访问"http://localhost:5050"将会看到文件
package main

import "net/http"

func main() {
	http.Handle("/", http.FileServer(http.Dir("./images")))
	http.ListenAndServe(":5050", nil)
}
