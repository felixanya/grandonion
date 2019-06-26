package server

import (
	"github.com/gin-gonic/gin"
)

type Server struct {

}

func (s *Server) Start() error {
	e := gin.Default()
	e.Use()

	vGroup := e.Group("/v1")
	a1 := NewApp1()
	//a2 := NewApp2()

	if a1 != nil {
		a1.Router(vGroup)
	}

	//if a2 != nil {
	//	a2.Router(vGroup)
	//}

	if err := e.Run(":8965"); err != nil {
		return err
	}
	return nil
}
