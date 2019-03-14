package http

import (
	"github.com/aaronize/grandonion/examples/http/controller/order"
	"github.com/aaronize/grandonion/examples/jobque"
	"github.com/gin-gonic/gin"
)


type Server struct {
	Listen 			string
	MaxWorkerNum 	int

}

func NewServer(path string) *Server {
	// path 配置文件路径
	// loadConf
	//
	listen := "localhost:9900"
	return &Server{
		Listen: listen,
		MaxWorkerNum: 10,
	}
}

// 初始化
func (s *Server) initWorkerQue() error {

	dispatcher := jobque.NewDispatcher(s.MaxWorkerNum)
	dispatcher.Run()

	return nil
}

func (s *Server) initRouter() error {
	r := gin.Default()

	orderGroup := r.Group("/order")
	proGroup := r.Group("/processor")
	//vmRouter(r.Group("/vm"))
	//resourceRouter(r.Group("/resource"))

	order.OrderRouter(orderGroup)
	processorRouter(proGroup)

	err := r.Run(s.Listen)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) RunServer() error {
	// path
	path := ""
	server := NewServer(path)
	// 初始化
	server.initWorkerQue()
	server.initRouter()

	return nil
}
