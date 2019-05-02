package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func UseXORM() {

	var err error
	engine, err = xorm.NewEngine("mysql", "root:123456@localhost:3306/testdb?charset=utf8")
	if err != nil {
		fmt.Println("new xorm engine error.", err.Error())
		return
	}
	fmt.Println("connect db success")


}

func main() {
	UseXORM()
}


