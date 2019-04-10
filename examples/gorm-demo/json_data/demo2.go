package main

import (
	"fmt"
	"github.com/aaronize/grandonion/examples/gorm-demo/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"strings"
)

const (
	userName1 = "root"
	password1 = "*"
	ip1       = "*"
	port1     = "3306"
	dbName1   = "*"
)

func getVIP(c *gin.Context) {

}

func setVIP(c *gin.Context) {
	vip := model.VIPerson{}
	if err := c.BindJSON(&vip); err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "failed", "data": err.Error(), "code": 1})
		return
	}
	fmt.Println(">>>check request parameter>>>", vip)
}

func main() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName1, ":", password1, "@tcp(", ip1, ":", port1, ")/", dbName1, "?charset=utf8"}, "")
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
