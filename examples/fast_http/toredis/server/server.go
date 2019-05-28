package server

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/go-redis/redis"
	"github.com/valyala/fasthttp"
	"log"
)

type Server struct {
	Config 		*Config
	client 		redis.ClusterClient
}

func (s *Server) Start() error {

	if err := s.Config.Logger.initLogger(); err != nil {
		//
		log.Fatal("init logger error.", err.Error())
	}

	if err := s.Config.Redis.initRedisCluster(); err != nil {
		//
		log.Fatal("init redis cluster error.", err.Error())
	}

	r, err := s.initRouter()
	if err != nil {
		log.Fatal("init router error.", err.Error())
	}

	if err := fasthttp.ListenAndServe(s.Config.Listen, r.Handler); err != nil {
		//
		log.Fatal("start http server error.", err.Error())
	}

	return nil
}

func (s *Server) initRouter() (*fasthttprouter.Router, error) {
	router := fasthttprouter.New()
	router.GET("/testreq", func(ctx *fasthttp.RequestCtx) {
		fmt.Fprintf(ctx, "Welcome!")
	})

	router.POST("/up", uploadToRedis)

	return router, nil
}
