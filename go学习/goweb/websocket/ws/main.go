package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"html/template"
	"net/http"
	"time"
)

//用户名称和信息
type userInfo struct {
	 Name string
	 Time string
}
//用户信息
type UserMsg struct {
	 Name string
	 Info string
}



//用户
var Usermap=make(map[string]userInfo)

var UserMsgList []UserMsg

var ZongCon=make(map[*websocket.Conn]bool)


//CheckOrigin防止跨站点的请求伪造
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main()  {
	http.HandleFunc("/login",login)
	http.HandleFunc("/dating",dating)
	http.HandleFunc("/visit",visit)
	http.ListenAndServe(":55", nil)

}
//登录页面
func login(w http.ResponseWriter, r *http.Request)  {
	t:=template.Must((template.ParseFiles("view/page/login.html")))
	t.Execute(w,"")
}

//大厅页面
func dating(w http.ResponseWriter, r *http.Request)  {
	//msglist:=make(map[int][]userMsg)
	mingzi:= string(r.URL.Query().Get("mingzi"))
	timeStr:=time.Now().Format("2006-01-02 15:04:05")
	Usermap[mingzi]=userInfo{Name:mingzi,Time:timeStr}
	type dataS struct {
		 NowName string
		 Usermap interface{}
		UserMsgList interface{}
 	}

	realdata:= dataS{NowName:mingzi,Usermap:Usermap,UserMsgList:UserMsgList}
	t:=template.Must((template.ParseFiles("view/page/dating.html")))
	t.Execute(w,realdata)
}

func  visit(w http.ResponseWriter, r *http.Request)  {
	conn, _ := upGrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
	ZongCon[conn]=true

   for{
	   // Read message from browser
	   msgType, msg, err := conn.ReadMessage()
	   if err != nil {  //读不到当前数据 相当于用户离开了
		   delete(ZongCon,conn)
		   fmt.Println("开始")
		   fmt.Println(err)
		   return
	   }
	   var tmpUserInfo UserMsg
	   json.Unmarshal([]byte(msg), &tmpUserInfo)
          fmt.Println(tmpUserInfo)
	  // tmpUserMsg=UserMsg{Name: tmpUserInfo.name,Info: tmpUserInfo.info}

	   UserMsgList=append(UserMsgList,tmpUserInfo)
	   // Print the message to the console
	   fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

	   // Write message back to browser
	   for tmpCon:=range  ZongCon{
		   if err = tmpCon.WriteMessage(msgType, msg); err != nil {
			   fmt.Println("发消息")
			   fmt.Println(err)
			   return
		   }
	   }





	   //if err = conn.WriteMessage(msgType, msg); err != nil {
		// return
	   //}
   }
}


