package dao

import (
	"wang/model"
	"wang/utils"
)

func CheckUserNameAndPassword(username string,password string)  (*model.User,error) {
	 sql:="select id,username,password,email from user where username=? and password=?"
      row:=utils.Db.QueryRow(sql,username,password)
	  user:=&model.User{}
	  row.Scan(&user.ID,&user.Username,&user.Password,&user.Email)
	  return user,nil
}
