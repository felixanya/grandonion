package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"strings"
)

var db *gorm.DB

const (
	userName = "root"
	password = "*"
	ip       = "*"
	port     = "3306"
	dbName   = "*"
)

type Iplist_iplist struct {
	Master_ip string `gorm:"type:varchar(100)"`
	Slave_ip  string `gorm:"type:varchar(100)"`
	Id        int    `gorm:"primary_key"`
	Location  string `gorm:"type:varchar(100)"`
	Latitude  string `gorm:"type:varchar(100)"`
	Longitude string `gorm:"type:varchar(100)"`
	Info 	map[string]string `gorm:"type:json"`
}

//注意方法名大写，就是public
//func ConnectionDB()  {
func main() {
	var err error
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	db, err = gorm.Open("mysql", path)

	//验证连接
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("connnect success")
	defer db.Close()
	r := gin.Default()
	err = db.AutoMigrate(&Iplist_iplist{}).Error
	if err != nil {
		fmt.Println("automigrate err")
	} else {
		fmt.Println("automigrate success")
	}
	r.GET("/", GetIP)
	r.POST("/", SetIP)
	r.Run(":8080")
}

func SetIP(c *gin.Context) {
	var IPInfo Iplist_iplist
	if err := c.BindJSON(&IPInfo); err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "failed", "data": err.Error(), "code": 1})
		return
	}
	fmt.Println(">>>check ipinfo>>>", IPInfo)
	if err := db.Create(&IPInfo).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "failed", "data": err.Error(), "code": 2})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": "", "code": 0})
	return
}

func GetIP(c *gin.Context) {
	var result []Iplist_iplist

	if err := db.Find(&result).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, result)
	}
}
