package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func DatabaseConnect() {
	//数据库连接
	db, _ := sql.Open("mysql", "ld:199935@(127.0.0.1:3306)/mydatabase?charset=utf8&parseTime=True&loc=Local")
	err := db.Ping()
	if err != nil {
		fmt.Println("数据库链接失败")
	}
	defer db.Close()
	//多行查询
	rows, _ := db.Query("select * from student")
	var name string
	var sco, age, gender, clazz int
	for rows.Next() {
		rows.Scan(&sco, &name, &age, &gender, &clazz)
		fmt.Println(sco, name, age, gender, clazz)
	}
}
