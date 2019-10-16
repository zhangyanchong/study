package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"bufio"

	"net/http"

	"golang.org/x/net/websocket"
)

type Room struct {
	Id       int    //群id
	RoomName string //群名
	MasterId int    //群主id
}
type Man struct {
	Id       int             //id
	Name     string          `json:"name"`     //昵称
	PassWord string          `json:"password"` //密码
	Friends  []string        //好友
	Grouds   []string        //组群
	Conn     *websocket.Conn `json:"-` //连接
}
type MessageData struct {
	Message  string //消息内容
	UserName string //发送者昵称
	ToUser   string //接受的用户id
	ToRoom   string //接受的群id，touserid和toroomid只存在一个
}

//建立连接阶段不知道连接的用户信息，在此临时储存连接
// var Conns []*websocket.Conn

var Rooms map[string]*Room
var Mans map[string]*Man

func main() {
	Mans = make(map[string]*Man, 0)
	Rooms = make(map[string]*Room, 0)

	// 初始化数据
	var room1 = Room{Id: 1, RoomName: "快乐一家人", MasterId: 1}
	Rooms[room1.RoomName] = &room1

	var man1 = Man{Id: 1, Name: "小明", PassWord: "111111"}
	man1.Grouds = append(man1.Grouds, "快乐一家人")
	man1.Friends = append(man1.Friends, "小红")
	Mans[man1.Name] = &man1

	var man2 = Man{Id: 2, Name: "小红", PassWord: "222222"}
	man2.Grouds = append(man2.Grouds, "快乐一家人")
	man2.Friends = append(man2.Friends, "小明")
	Mans[man2.Name] = &man2

	http.Handle("/chatroom", websocket.Handler(ChatroomServer)) //监听ws连接
	http.HandleFunc("/login", login)                            //监听登陆
	// http.HandleFunc("/", Client)
	err := http.ListenAndServe(":6611", nil)

	if err != nil {

		panic("ListenAndServe: " + err.Error())

	}
}

func ChatroomServer(ws *websocket.Conn) {

	fmt.Printf("connection a man \n")
	conneduser := ""
	r := bufio.NewReader(ws)
	for {
		data, err := r.ReadBytes('\n')
		if err != nil {
			fmt.Printf("disconnected \n")
			var data = MessageData{
				Message:  "you disconnected",
				UserName: "sys",
			}
			SenToUser(ws, data)
			break
		}

		var message MessageData
		err = json.Unmarshal(data, &message)
		if err != nil {
			fmt.Printf("%s \n", err)
		}
		if message.Message == "[*login*]" {
			v, ok := Mans[message.UserName]
			if ok { //有此人
				if v.Conn == nil { //没连接过
					v.Conn = ws
					if v.Conn == nil {
						fmt.Printf("777 \n")
					}
					conneduser = v.Name
					fmt.Printf("%s login success \n", message.UserName)
					var data = MessageData{
						Message:  "login success",
						UserName: "sys",
					}
					SenToUser(ws, data)
					continue
				} else {
					fmt.Printf("%s try login,but logined \n", message.UserName)
					var data = MessageData{
						Message:  "login fail",
						UserName: "sys",
					}
					SenToUser(ws, data)
					continue
				}

			} else {
				fmt.Printf("no user %s \n", message.UserName)
				var data = MessageData{
					Message:  "login fail",
					UserName: "sys",
				}
				SenToUser(ws, data)
				continue
			}
		} else {
			if v, ok := Mans[message.UserName]; ok {
				if v.Conn != nil {

					SendMessage(ws, message)
				}

			}

		}
	}
	defer func() {
		fmt.Printf("%s 关闭了连接 \n", conneduser)
		if conneduser != "" {
			v, _ := Mans[conneduser]
			v.Conn = nil
		}
		ws.Close()
	}()
}

//发送信息给用户
func SenToUser(ws *websocket.Conn, data MessageData) {
	message, _ := json.Marshal(data)
	// io.WriteString(ws, )
	// fmt.Fprintf(ws, "%s", message)
	// message, _ := json.Marshal(data)
	_, err := io.WriteString(ws, string(message))
	if err != nil {
		fmt.Printf("%s \n", err)
	}
}
func SendMessage(self *websocket.Conn, data MessageData) {
	if data.ToUser != "0" { //发给用户
		if v, ok := Mans[data.ToUser]; ok && v.Conn != nil {
			for _, friend := range v.Friends {
				if data.UserName == friend { //判断是否是好友
					fmt.Printf("%s to %s: %s \n", data.UserName, data.ToUser, data.Message)
					SenToUser(self, data) //给自己发一份
					SenToUser(v.Conn, data)
					break
				}
			}

		}
		return
	}
	if data.ToRoom != "0" { //发给组群
		fmt.Printf("%s to %s: %s \n", data.UserName, data.ToRoom, data.Message)
		for _, v := range Mans {
			if v.Conn != nil {
				for _, room := range v.Grouds {
					if data.ToRoom == room {
						message, _ := json.Marshal(data)

						_, err := io.WriteString(v.Conn, string(message))
						if err != nil {
							fmt.Printf("%s \n", err)
						}
					}
				}

			}
		}
		return
	}

}
func login(w http.ResponseWriter, r *http.Request) {
	// fmt.Printf(r.Method)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type") //因为前端请求头包含了自定义的字段“content-type”，所以这里要允许他出现
	if r.Method != "OPTIONS" {
		s, _ := ioutil.ReadAll(r.Body)
		var user Man
		err := json.Unmarshal(s, &user)
		if err != nil {
			fmt.Printf("%s", err)
		}
		// fmt.Printf(" %v %v", user.Name, user.PassWord)
		if v, ok := Mans[user.Name]; ok && v.PassWord == user.PassWord {
			str, _ := json.Marshal(v)
			// fmt.Printf("%s", str)
			fmt.Fprintf(w, "%s", str)
		} else {
			fmt.Fprintf(w, "%s", "null")
		}
	}

}
