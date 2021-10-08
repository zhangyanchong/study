package api

import "github.com/gogf/gf/net/ghttp"


type IndexApi struct{}

var Index = new(IndexApi)

func (c *IndexApi) Index(r *ghttp.Request) {
	r.Response.Write("api接口使用")
}