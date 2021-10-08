package model

type AccountReqLoign struct{
    Username  string  `p:"username"  v:"required#用户名称不能为空"`
    PassWord  string  `p:"password"  v:"required#密码不能为空"`
}

//用户信息 userinfo
type UserInfo struct {
    Id           int `json:"id"`     //
    Username     string `json:"username"`
    Password     string `json:"password"`
    NickName     string  `json:"nickName"`
}


