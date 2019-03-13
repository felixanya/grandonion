package http

import (
	"fmt"
	"github.com/aaronize/grandonion/examples/http/test"
	"github.com/aaronize/grandonion/examples/jobque"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var (
	DB 	*gorm.DB
	JobQueue 	chan Job
)

type Job struct {
	State 		string
	Action 		func([]string) error
	UID 		string
	Parm 		[]string

	Success		chan string
	Fail 		chan error
	Finish 		chan bool
}


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
	// 初始化Job
	JobQueue = make(chan Job, 1000)

	dispatcher := jobque.NewDispatcher(s.MaxWorkerNum)
	dispatcher.Run()

	return nil
}

func initRouter(engine *gin.Engine) *gin.Engine {
	// 添加中间件
	engine.Use()

	// 开始路径
	uperRoute := engine.Group("/v1")
	RouteDispatcher(uperRoute)

	testRoute := engine.Group("/test")
	test.TestDispatcher(testRoute)

	return engine
}

func (s *Server) RunServer() error {
	// path
	path := ""
	server := NewServer(path)
	// 初始化
	server.initWorkerQue()

	r := initRouter(gin.Default())
	if err := r.Run(s.Listen); err != nil {
		fmt.Println("Run server error", err.Error())
		return err
	}

	return nil
}
