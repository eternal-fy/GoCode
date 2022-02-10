package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"strconv"
)

type Person struct {
	gorm.Model
	Age  *int32
	Name string
}

//钩子函数，在添加新数据前设置列名的值
func (P *Person) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("age", 999)
	return nil
}

func (person Person) String() string {
	return "Person{" + "age:" + strconv.Itoa(int(*(person.Age))) + " Name：" + person.Name + "}"
}

func (*Person) TableName() string {
	return "Myperson"
}

func Handle() {
	db, err := gorm.Open("mysql", "ld:199935@(127.0.0.1:3306)/mydatabase?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		panic("链接数据库失败")
	}
	createTable(db)
	var persons []Person
	db.Limit(1).Find(&persons)
	for _, person := range persons {
		fmt.Println(person)
	}
	var age int32
	age = 10
	p := Person{Age: &age, Name: "dd"}
	fmt.Println(p)
	db.Create(&p)
	//}

	//p := &Person{Name: "ld"} //如果age字段是int，则会写入age为0，为了避免填值，设置为*int32
	//pack := make([]Person, 10)
	//db.Create(p)             //添加
	//db.Table("myperson").Update(p).Where("name")
	//查询第一个
	//db.Table("myperson").First(&pack)
	//fmt.Printf("%v", pack)

	////查询所有
	//db.Find(&pack)
	//fmt.Printf("%v", pack)

	//where查询
	//db.Where("id = ?", 3).Find(&pack)
	//fmt.Printf("%v", pack)

	//not in
	//db.Not("name", "jinzhu").First(&user)
}

func createTable(db *gorm.DB) {
	db.AutoMigrate(&Person{}) //如果数据库表中缺少该结构体的某个字段，则补齐。

}
