package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

var DB *gorm.DB

type Gorm struct{

}



func (Student) TableName() string {
	fmt.Println("创建的时候会去寻找这个方法")
	return "student"
}
func (this *Gorm) Init() {
	var err error
	conn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		Mysqlconn["username"],
		Mysqlconn["password"],
		Mysqlconn["host"],
		Mysqlconn["port"],
		Mysqlconn["name"],
	)
	fmt.Print("mysqlconn:")
	fmt.Println(Mysqlconn)
	DB, err = gorm.Open("mysql", conn)
	if err != nil {
		panic(err.Error())
	}
	if DB.HasTable("student") {
	fmt.Println("数据表已经存在")
	} else {
		DB.CreateTable(&Student{})
	}

}


