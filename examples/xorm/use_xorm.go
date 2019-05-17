package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
)

var engine *xorm.Engine

type Switches struct {
	Id 			int 		`json:"id" xorm:"'id' pk autoincr"`
	CreatedAt 	time.Time	`json:"created_at" xorm:"'created_at' created"`
	UpdatedAt 	time.Time	`json:"updated_at" xorm:"'updated_at' updated"`
	Ip 			string		`json:"ip" xorm:"'ip'"`
	HostName 	string		`json:"host_name" xorm:"'host_name'"`
	State		string		`json:"state" xorm:"'state'"`
	Info 		string		`json:"info" xorm:"'info'"`
	Version 	 int 		`xorm:"'version' version"`
}

func UseXORM() {

	var err error
	engine, err = xorm.NewEngine("mysql", "root:123456@/testdb?charset=utf8")
	if err != nil {
		fmt.Println("new xorm engine error.", err.Error())
		return
	}
	fmt.Println("connect db success")

	if err := engine.Sync2(new(Switches)); err != nil {
		panic(err)
	}
	var switcher Switches
	var has bool
	if has, err = engine.Id(1).Get(&switcher); err != nil {
		panic(err)
	}
	if !has {
		panic("未找到数据")
	}
	fmt.Println(switcher)

	var switcher2 Switches
	switcher2.Info = "demo1测试数据4"
	switcher2.Version = 3
	_, err = engine.ID(1).Update(&switcher2)
	if err != nil {
		panic(err)
	}
}

func main() {
	UseXORM()
}


