package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

//默认表名为users   如果结构体名称中有大写（除了首字母），会通过_隔开，例如  struct  myHome  对应数据库表名my_home
type User struct {
	gorm.Model
	Name     string    `gorm:"type:varchar(255)"`
	Age      int       `gorm:"type:int"`
	Birthday time.Time `gorm:"type:datetime"`
}

/*设置 `User` 的表名为 `profiles`
func (User) TableName() string {
	return "profiles"
}*/

func main() {
	db := ConnectMysql()
	defer db.Close()
	db.Table("users").Where("age=?", "22").Update("age", 10)

	//根据结构体创建数据库
	//db.Table("users").CreateTable(&User{})
	user := User{Name: "aa", Age: 22, Birthday: time.Now()}
	flag1 := db.NewRecord(user) // => 返回 `true` ，因为主键为空
	println(flag1)
	db.Create(&user)

	flag2 := db.NewRecord(user) // => 在 `user` 之后创建返回 `false`
	println(flag2)
}
func ConnectMysql() *gorm.DB {
	db, err := gorm.Open("mysql", "ld:199935@(localhost:3306)/mydatabase?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	return db
}

/*
gorm 字段类型设置
type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"` //设置字段的大小为255个字节
	MemberNumber *string `gorm:"unique;not null"` // 设置 memberNumber 字段唯一且不为空
	Num          int     `gorm:"AUTO_INCREMENT"` // 设置 Num字段自增
	Address      string  `gorm:"index:addr"` // 给Address 创建一个名字是  `addr`的索引
	IgnoreMe     int     `gorm:"-"` //忽略这个字段
}*/
