package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/test", func(c *gin.Context) {
		param := struct {
			Action 	string 	`json:"action" binding:"required"`
			UID 	string	`json:"uid" binding:"required"`
			Required 	map[string]interface{}	`json:"required"`
			Optional 	map[string]interface{}	`json:"optional"`
		}{}

		if err := c.ShouldBindJSON(&param); err != nil {
			fmt.Println("bind json error:", err.Error())
			return
		}

		fmt.Println(param.Required["owner"])
		for i := range param.Required {
			fmt.Println("param: ", i)
		}
		return
	})

	r.Run(":9900")
}
