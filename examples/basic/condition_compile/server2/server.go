package server2

import (
	"github.com/gin-gonic/gin"
)

// 条件编译格式要求严格：
// 1. // +build下要空一行
// 2. 条件间","表示与，" "(空格)表示或，换行表示与

// 这个例子的编译逻辑：
// go run main.go // 默认情况编译全部
// go run -tags="app1" main.go // 编译app1部分
// go run -tags="app2" main.go // 编译app2部分
// go run -tags="app1 app2" main.go // 编译app1和app2部分

// 当判断的条件多了会比较麻烦

type Server struct {

}

func (s *Server) Start() error {
	e := gin.Default()
	e.Use()

	vGroup := e.Group("/v1")
	a1 := NewApp1()
	a2 := NewApp2()

	if a1 != nil {
		a1.Router(vGroup)
	}
	if a2 != nil {
		a2.Router(vGroup)
	}

	if err := e.Run(":8965"); err != nil {
		return err
	}
	return nil
}
