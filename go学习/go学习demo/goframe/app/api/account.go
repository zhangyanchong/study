package api

import (
	"github.com/gogf/gf/net/ghttp"
	"xin/app/model"
	"xin/app/service"
	"xin/app/utils/helper"
)


type AccountApi struct{}

var Account = new(AccountApi)

func (c *AccountApi) Login(r *ghttp.Request) {
	var req *model.AccountReqLoign
	//验证错误
	if err := r.Parse(&req); err != nil {
		helper.Help.ReturnJson(r,201,err.Error(),"")
	}
     userInfo:=service.Accout.UserQuery(req)
	 if(userInfo!=nil){ //登录成功
		 r.Session.Set("userId",userInfo.Id)
		 r.Session.Set("username",userInfo.Username)
		 r.Session.Set("nickName",userInfo.NickName)
		 helper.Help.ReturnJson(r,200,"ok","")
	 }else{  //登录失败
		 helper.Help.ReturnJson(r,201,"fail","")
	 }
}

/*用户信息**/
func (c *AccountApi)UserInfo(r *ghttp.Request)  {
    userId:=r.Session.GetInt64("userId");
	userInfo:=service.Accout.UserIdQuery(userId)
	helper.Help.ReturnJson(r,200,"ok",userInfo)
}

/*添加用户**/
func  (c *AccountApi)AddUser(r *ghttp.Request)  {
	username:=r.GetRequestString("username")
	password:=r.GetRequestString("password")
	if(username==""){
		helper.Help.ReturnJson(r,201,"用户名不能为空","")
	}
	insertRs:=service.Accout.AddUser(username,password)
	if(insertRs!=0){
		helper.Help.ReturnJson(r,200,"ok","")
	}else{
		helper.Help.ReturnJson(r,201,"fail","")
	}

}

/*修改用户**/
func  (c *AccountApi)UpdateUser(r *ghttp.Request)  {
	userId:=r.GetRequestInt64("id")
	username:=r.GetRequestString("username")
	if(userId==0){
		helper.Help.ReturnJson(r,201,"用户id不能为空","")
	}
	userInfo:=service.Accout.UserIdQuery(userId)
	if(userInfo==nil){
		helper.Help.ReturnJson(r,201,"用户信息不存在","")
	}
	updateRs:=service.Accout.UpdateUser(userId,username)
	if(updateRs==true){
		helper.Help.ReturnJson(r,200,"ok","")
	}else{
		helper.Help.ReturnJson(r,201,"fail","")
	}
}

/*用户列表**/
func  (c *AccountApi)UserList(r *ghttp.Request)  {
	username:=r.GetRequestString("username")
	nickname:=r.GetRequestString("nickname")
	rs:=service.Accout.UserList(username,nickname)
	helper.Help.ReturnJson(r,200,"ok",rs)
}


/*退出登录**/
func  (c *AccountApi)LoginOut(r *ghttp.Request)  {
	  r.Session.Clear()
	helper.Help.ReturnJson(r,200,"ok","")
}

func  (c *AccountApi)DelUser(r *ghttp.Request)  {
	userId:=r.GetRequestInt64("id")
	rs:=service.Accout.DelUser(userId)
	if(rs==true){
		helper.Help.ReturnJson(r,200,"ok","")
	}else{
		helper.Help.ReturnJson(r,201,"fail","")
	}
}