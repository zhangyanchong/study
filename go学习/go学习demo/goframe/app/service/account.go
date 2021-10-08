package service

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"xin/app/model"
)

// 中间件管理服务
var Accout = new(accountService)

type accountService struct{}

// 用户查询
func (s *accountService)UserQuery(r *model.AccountReqLoign)  *model.UserInfo{
   var	userInfo *model.UserInfo
	g.Model("user").Where(g.Map{"username":r.Username,"password":r.PassWord}).Scan(&userInfo)
   return  userInfo
}

//用户id 查询
func (s *accountService)UserIdQuery(id int64) *model.UserInfo  {
	var	userInfo *model.UserInfo
	g.Model("user").Where(g.Map{"id":id}).Scan(&userInfo)
	return  userInfo
}

//添加用户
func (s *accountService)AddUser(username string,password string)  int64  {
   lastId,_:=g.Model("user").InsertAndGetId(g.Map{"username": username,"password":password,"create_time":gtime.Timestamp(),"update_time":gtime.Timestamp()})
   return lastId
}

//修改用户
func (s *accountService)UpdateUser(id int64,username string)  bool {
	_,error:=g.Model("user").Data(g.Map{"username" : username}).Where("id", id).Update()
	if(error!=nil){
		return  false
	}
	return  true
}

//用户列表
func (s *accountService)UserList(username string,nickname string) []*model.UserInfo {
	query:=g.Model("user");
	if(username!=""){
		query=query.Where(g.Map{"username" : username})
	}
	if(nickname!=""){
		query=query.Where(g.Map{"nickname" : nickname})
	}
	var	userInfo []*model.UserInfo
	query.Scan(&userInfo)
   return  userInfo
}

func (s *accountService)DelUser(id int64) bool {
	_,error:=g.Model("user").Where("id", id).Delete()
	if(error!=nil){
		return  false
	}
	return  true
}