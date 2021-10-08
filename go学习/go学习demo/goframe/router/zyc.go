package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"xin/app/api"
	"xin/app/api/middleware"
)

func  init()  {
	s := g.Server()

	s.Group("api", func(group *ghttp.RouterGroup) {
		group.ALL("/index/index", api.Index.Index)
	    group.ALL("/account/login", api.Account.Login)
	})

	/* 中间件 */
	s.Group("api", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.LoginMid.Login)
		group.GET("/account/userInfo", api.Account.UserInfo)
		group.ALL("/account/loginOut", api.Account.LoginOut)
		group.ALL("/account/addUser", api.Account.AddUser)
		group.ALL("/account/updateUser", api.Account.UpdateUser)
		group.ALL("/account/userlist", api.Account.UserList)
		group.ALL("/account/delUser", api.Account.DelUser)
	})
}
