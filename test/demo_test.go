package test

import (
	"encoding/binary"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"strconv"
	"testing"
)

type Person struct {
	gorm.Model
	Age  *int32
	Name string
}

func (person Person) String() string {
	return "Person{" + "age:" + strconv.Itoa(int(*(person.Age))) + " Name：" + person.Name + "}"
}

func (*Person) TableName() string {
	return "Myperson"
}
func TestDemo(t *testing.T) {
	buf := []byte{'0', '1', '2'}
	data := make([]byte, 10)
	copy(data, buf)
	println(string(data))
	//io.ReadFull()
	//binary.BigEndian.PutUint32()
}

func TestDemo2(t *testing.T) {
	var data []byte = []byte("你好")
	buf := make([]byte, 4+len(data))
	binary.BigEndian.PutUint32(buf[:4], uint32(len(data))) //将len（data）的结果放进buf前四个，比如6, 00000000 00000000 00000000 00000110   分成四个字节放进buf去。
	copy(buf[4:], data)
	fmt.Println(len(data))
	fmt.Printf("%v,%s    ", buf[:4], buf[4:])
}

type Student struct {
	Name string
	Age  int
}

func TestDB(t *testing.T) {
	db := GetConn()
	defer db.Close()
	var persons []Person
	db.Limit(999).Where("age = ?", 999).Find(&persons)
	for _, person := range persons {
		fmt.Println(person)
	}

}
func TestStudent(t *testing.T) {
	db := GetConn()
	defer db.Close()
	//var students []Student
	var students []Student
	var count int
	db.Table("student").Select("name,age").Order("age").Offset(5).Limit(10).Find(&students)
	//db.Table("student").Select("name,age").Where("age = ?", 1).Find(&students).Count(&count)
	db.Table("student").Select("count(*)").Where("age != ?", 1).Count(&count)
	for _, student := range students {
		fmt.Println(student)
	}
	fmt.Println(count)
}

func TestGroup(t *testing.T) {
	db := GetConn()
	defer db.Close()
	//var students []Student
	var students []Student
	var count int
	db.Table("student").Where("age < ?", 15).Update("name", "刘德")
	rows, _ := db.Table("student").Select("name,SUM(age)").Group("name").Having("SUM(age) < ?", 100).Rows()
	for _, student := range students {
		fmt.Println(student)
	}
	for rows.Next() {
		var (
			name string
			sum  int
		)
		rows.Scan(&name, &sum)
		println(name, sum)
	}

	fmt.Println(count)
}

func GetConn() *gorm.DB {
	db, err := gorm.Open("mysql", "ld:199935@(127.0.0.1:3306)/mydatabase?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("链接数据库失败")
	}
	return db
}
