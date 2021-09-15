package dao

import (
	"fmt"
	"testing"
)

func TestCheckUserNameAndPassword(t *testing.T)  {
   user,_:=	CheckUserNameAndPassword("admin","admin")
   fmt.Println("user=",user)
}
