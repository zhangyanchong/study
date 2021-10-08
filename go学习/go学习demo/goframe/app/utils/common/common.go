package common

type JsonResult struct {
	Code  int         `json:"code"`  // 响应编码：200成功 401请登录 403无权限 500错误
	Msg   string      `json:"msg"`   // 消息提示语
	Data  interface{} `json:"data"`  // 数据对象
}

