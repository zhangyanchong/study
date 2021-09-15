package dao

import (
	"wang/model"
	"wang/utils"
)

func  GetBook()([] *model.Book,error) {
    sql:="select  id,title,author,price from book "
    rows,err:= utils.Db.Query(sql)
	if(err!=nil){
			return  nil,err
	}
	var books []*model.Book
	for rows.Next(){
		book:=&model.Book{}
		rows.Scan(&book.ID,&book.Title,&book.Author,&book.Price)
		books=append(books,book)
	}
	return books,nil
}
