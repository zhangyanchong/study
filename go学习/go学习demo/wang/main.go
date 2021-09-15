package main

import (
	"html/template"
	"net/http"
	"wang/controller"
)

func Indexhandler(w http.ResponseWriter, r *http.Request){
  t:=template.Must((template.ParseFiles("views/index.html")))

    type  infoS	struct {
         Name string
		 Add  string
	}
	info:=infoS{Name: "张三",Add: "ting"}
  t.Execute(w,info)
}

func Loginhandler(w http.ResponseWriter, r *http.Request)  {
	
}

func main()  {
	//处理静态文件
	http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("views/static/"))))
	http.Handle("/page/",http.StripPrefix("/page/",http.FileServer(http.Dir("views/page/"))))

	//首页
	http.HandleFunc("/main",Indexhandler)
   //登录
	http.HandleFunc("/login",controller.Login)

  //列表
    http.HandleFunc("/list",controller.BookList)

	//普通json 返回
	http.HandleFunc("/json",controller.Rjson)

	//列表json
	http.HandleFunc("/jsonlist",controller.JsonList)


	http.ListenAndServe(":8080",nil)
}
