package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/?charset=utf8")//对应数据库的用户名和密码
	defer db.Close()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("success")
	}
	db.Query("DROP database if EXISTS mydb")
	db.Query("CREATE database mydb")
	db.Query("CREATE TABLE mydb.squarenum(number INT,name VARCHAR(20))")
	db.Query("INSERT INTO mydb.squarenum VALUES (22,'shibiao1'),(11,'shibiao2'),(33,'shibiao3')")

	rows, err := db.Query("SELECT * FROM mydb.squarenum")
	checkError(err)
	for rows.Next() {
		var number int
		var name string
		err = rows.Scan(&number,&name)
		if err != nil {
			fmt.Println("error   :")
			panic(err)
		}
		fmt.Println(number,name)
	}
}
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}