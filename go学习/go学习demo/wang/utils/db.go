package utils

import "database/sql"
import 	_ "github.com/go-sql-driver/mysql"

var (
	Db *sql.DB
	err error
)

func init()  {
	Db,err=sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/bookstore")
	if(err!=nil){
		panic(err.Error())
	}
}