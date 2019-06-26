// +build !app1,app2

package server2

import (
	"github.com/gin-gonic/gin"
	"log"
)

type App1 struct {

}

func NewApp1() *App1 {
	log.Println(">>>in NewApp1 return nil>>>")
	return nil
}

func (a1 *App1) Router(rg *gin.RouterGroup) {

}
