package main

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	userName = ""
	password = ""
	ip       = ""
	port     = "3306"
	dbName   = "tmp"
)

func Dbcon() (db *sql.DB, errstr error) {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	db, db_error := sql.Open("mysql", path)

	//验证连接
	if err := db.Ping(); err != nil {
		fmt.Println("opon database fail")

	}
	return db, db_error

}
func main() {
	del()
}

//添加
func insertDb() {
	dbcon, _ := Dbcon()
	fmt.Println("Hello, World!")
	stmt, err := dbcon.Prepare("INSERT INTO user (`name`, `sex`,`createTime`) VALUES (?, ?,?)")
	if err != nil {
		fmt.Println("Prepare fail")
	}
	timestamp := time.Now().Unix()
	//将参数传递到sql语句中并且执行
	rs, err := stmt.Exec("zhangsan", "1", timestamp)
	if err != nil {
		fmt.Println("Exec fail")
	}
	id, err := rs.LastInsertId()
	fmt.Println("id:", id)
}

//修改
func updateDb() bool {
	tx, err := Dbcon()
	if err != nil {
		fmt.Println("tx fail")
	}
	//准备sql语句
	stmt, err := tx.Prepare("UPDATE user SET name = ?, sex = ? WHERE id = ?")
	if err != nil {
		fmt.Println("Prepare fail")
		return false
	}
	//设置参数以及执行sql语句
	res, err := stmt.Exec("hahha", 4, 3)
	if err != nil {
		fmt.Println("Exec fail")
		return false
	}
	res.LastInsertId()
	return true
}

type tmpUser struct {
	id         string
	name       string
	sex        string
	createTime string
}

//查找
func cha() {
	db, err := Dbcon()
	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	for rows.Next() {
		var student tmpUser
		err = rows.Scan(&student.id, &student.name, &student.sex, &student.createTime)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		fmt.Println(student)
	}

}

//删除
func del() bool {
	db, err := Dbcon()
	res, err := db.Exec("delete from user where id=?", 1)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return false
	}
	res.LastInsertId()
	return true
}
