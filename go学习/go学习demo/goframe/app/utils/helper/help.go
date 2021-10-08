package helper

import (
	"github.com/gogf/gf/net/ghttp"
	"xin/app/utils/common"
)

var Help = new(helpHelper)

type helpHelper struct{}

func (h helpHelper) ReturnJson(r *ghttp.Request,code int ,msg string,data interface{} )  {
      json:=common.JsonResult{Code: code,Msg: msg,Data:data}
	  r.Response.WriteJsonExit(json)
}
