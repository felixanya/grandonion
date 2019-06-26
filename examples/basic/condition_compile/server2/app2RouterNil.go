// +build !app2,app1

package server2

import (
	"github.com/gin-gonic/gin"
	"log"
)

type App2 struct {

}

func NewApp2() *App2 {
	log.Println(">>>in NewApp2 return nil>>>")
	return nil
}

func (a2 App2) Router(rg *gin.RouterGroup) {

}
