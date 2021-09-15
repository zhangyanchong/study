package controller

import (
	"html/template"
	"net/http"
	"wang/dao"
)

func Login(w http.ResponseWriter, r *http.Request)  {
	username:=r.PostFormValue("username")
	password:=r.PostFormValue("password")
	user,_:=dao.CheckUserNameAndPassword(username,password)
	if user.ID>0{
       t:= template.Must(template.ParseFiles("views/page/login_success.html"))
	   t.Execute(w,"")
	}else{
		t:= template.Must(template.ParseFiles("views/page/login.html"))
		t.Execute(w,"")
	}
}
