package main

import (
	model2 "github.com/aaronize/grandonion/examples/db_orm/gorm-demo/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

var gdb *gorm.DB

func main() {
	defer gdb.Close()

	s1 := &model2.Server{
		IP: "192.168.32.128",
		HostName: "192-168-32-128",
		State: "running",
		Info: "",
		ServerTags: model2.ServerTags{
			map[string]interface{}{
				"Type": "type1",
				"Usage": "ut1",
				"Location": "sh,cn",
				"Tags": []string{"tag3"},
			},
		},
	}
	if err := gdb.Create(s1).Error; err != nil {
		log.Fatal("插入数据错误，", err.Error())
		return
	}

	var serverList []model2.Server
	if err := gdb.Find(&serverList).Error; err != nil {
		log.Fatal("查询数据库错误，", err.Error())
		return
	}
	for _, s := range serverList {
		log.Println("server: ", s.ServerTags)
	}
}

func init() {
	var err error
	// "user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
	gdb, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/testdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("连接数据库错误，", err.Error())
		os.Exit(-1)
	} else {
		log.Println("数据库连接成功。")
	}
	gdb.AutoMigrate(&model2.Server{})
}
