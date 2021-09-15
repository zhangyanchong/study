package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"wang/dao"
)

func BookList(w http.ResponseWriter, r *http.Request)  {
    type dataS	struct {
	   Num  int
	   Data  interface{}
	}

	booklist,_:= dao.GetBook()
	//for k,v:=range booklist{
    //   fmt.Println("值是：",k,v)
	//}
	 realdata:= dataS{Num:5,Data:booklist}
	t:=template.Must((template.ParseFiles("views/page/list.html")))
	t.Execute(w,realdata)
}

func Rjson(w http.ResponseWriter, r *http.Request)  {
	id:= r.URL.Query().Get("id")
	if(id==""){
		fmt.Fprint(w, "id不能为空")
		return
	}
	type info struct {
	     Code int
		 Msg string
		 XinData interface{}
	}
	realData:=info{Code: 1,Msg:"ok"}
	data, _ := json.Marshal(realData)
	fmt.Fprint(w, string(data))
}

func JsonList(w http.ResponseWriter, r *http.Request)  {
	xinid:=r.PostFormValue("id")
	if(xinid==""){
		println("id不能为空")
	}
	type dataS	struct {
		Num  int
		Data  interface{}
	}
	booklist,_:= dao.GetBook()

	realData:= dataS{Num:5,Data:booklist}
	newdata, _ := json.Marshal(realData)
	fmt.Fprint(w, string(newdata))
}