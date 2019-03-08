package processor

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Processor(c *gin.Context) {
	param := struct {
		Data 	string	`json:"data"`
	}{}
	if err := c.BindJSON(&param); err != nil {
		fmt.Println("参数错误,", err.Error())
		return
	}

	// 处理逻辑


	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"code": 0,
		"data": "上报成功",
	})
}
