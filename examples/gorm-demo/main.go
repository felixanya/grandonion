package main

import (
	"fmt"
	"github.com/aaronize/grandonion/examples/gorm-demo/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strings"
)
const (
	userName = "root"
	password = "Da!chen1124"
	ip       = "39.108.171.31"
	port     = "3306"
	dbName   = "gormdb"
)

func main() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	db, err := gorm.Open("mysql", path)

	//验证连接
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("connnect success")
	defer db.Close()
	r := gin.Default()
	err = db.AutoMigrate(&model.VIPerson{}).Error
	if err != nil {
		fmt.Println("automigrate err")
	} else {
		fmt.Println("automigrate success")
	}



	r.GET("/", getVIP)
	r.POST("/", setVIP)

	r.Run(":8080")
}
