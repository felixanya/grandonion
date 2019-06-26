// +build app2,!app1 !app2,!app1 app1,app2

package server2

import (
	"github.com/gin-gonic/gin"
	"log"
)

type App2 struct {

}

func NewApp2() *App2 {
	log.Println(">>>in NewApp2>>>")
	return &App2{}
}

func (a2 *App2) Router(rg *gin.RouterGroup) {
	rg.Use()

	rg.GET("/a2", a2.get)
	rg.POST("/a2", a2.post)
	rg.PUT("/a2", a2.put)
}

func (a2 *App2) get(c *gin.Context) {

}

func (a2 *App2) post(c *gin.Context) {

}

func (a2 *App2) put(c *gin.Context) {

}
