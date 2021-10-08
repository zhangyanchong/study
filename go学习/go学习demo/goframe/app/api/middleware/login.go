package middleware

import (
	"github.com/gogf/gf/net/ghttp"
	"xin/app/utils/helper"
)

type LoginMidServer struct{}

var  LoginMid = new(LoginMidServer)

//中间件使用
func (LoginMidServer)Login(r *ghttp.Request)  {
	userId:=r.Session.GetInt32("userId");
	if(userId==0){
		helper.Help.ReturnJson(r,-1,"need login","")
	}
	r.Middleware.Next()
}
