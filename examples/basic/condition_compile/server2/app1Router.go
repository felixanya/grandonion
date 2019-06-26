// +build app1,!app2 !app1,!app2 app1,app2

package server2

import (
	"github.com/gin-gonic/gin"
	"log"
)

type App1 struct {

}

func NewApp1() *App1 {
	log.Println(">>>in NewApp1>>>")
	return &App1{}
}

func (a1 *App1) Router(rg *gin.RouterGroup) {

	rg.Use()

	rg.GET("/a1", a1.get)
	rg.POST("/a1", a1.post)
	rg.PUT("/a1", a1.put)
}

func (a1 *App1) get(c *gin.Context) {

}

func (a1 *App1) post(c *gin.Context) {

}

func (a1 *App1) put(c *gin.Context) {

}
