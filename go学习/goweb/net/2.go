package main

//下载文件
import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// 获取文件
	imageUrl := "http://pic.uzzf.com/up/2015-7/20157816026.jpg"
	response, err := http.Get(imageUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	// 创建保存位置
	file, err2 := os.Create("pic.jpg")
	if err2 != nil {
		panic(err2)
	}
	// 保存文件
	_, err3 := io.Copy(file, response.Body)
	if err3 != nil {
		panic(err3)
	}

	file.Close()
	fmt.Println("Image downloading is successful.")
}
