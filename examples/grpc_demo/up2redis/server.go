package up2redis

import (
	"context"
	"fmt"
	pb "github.com/aaronize/grandonion/examples/grpc_demo/up2redis/toredis"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
	"time"
)

type Pool struct {
	queue  chan int
	wg 	*sync.WaitGroup
}

func New(size int) *Pool {
	if size <= 0 {
		size = 1
	}
	return &Pool{
		queue: make(chan int, size),
		wg:    &sync.WaitGroup{},
	}
}

func (p *Pool) Add(delta int) {
	for i := 0; i < delta; i++ {
		p.queue <- 1
	}
	for i := 0; i > delta; i-- {
		<-p.queue
	}
	p.wg.Add(delta)
}

func (p *Pool) Done() {
	<-p.queue
	p.wg.Done()
}

func (p *Pool) Wait() {
	p.wg.Wait()
}



type Server struct {
	Config 		*Config

}


func (s *Server) Start() error {
	err := s.Config.Logger.initLogger()
	if err != nil {
		log.Fatalf("Initial logger error. %s", err.Error())
	}

	err = s.Config.Redis.initRedisCluster()
	if err != nil {
		log.Fatalf("Initila redis cluster error. %s", err.Error())
	}

	li, err := net.Listen("tcp", s.Config.Listen)
	log.Println("Grpc listening ", s.Config.Listen)
	if err != nil {
		log.Fatalf("Failed to listen: %s", err.Error())
	}

	rs := grpc.NewServer()
	pb.RegisterRedisServer(rs, s)
	if err := rs.Serve(li); err != nil {
		log.Fatalf("Failed to serve, %s", err.Error())
	}
	return nil
}


func (s *Server) Report(ctx context.Context, in *pb.ReportRequest) (*pb.ReportResponse, error) {
	pool := New(s.Config.PoolSize)
	pool.Add(1)
	go s.uploader(in, pool)

	return &pb.ReportResponse{Info: "push data success"}, nil
}

func (s *Server) uploader(in *pb.ReportRequest, pool *Pool) {
	if in.Delay == 0.0 {
		err := RedisClient().HSet(in.Key, string(in.Timestamp), in.Delay).Err()
		if err != nil {
			logger.Error("push data failed", zap.String("detail", err.Error()),
				zap.String("key", in.Key))
			//return &pb.ReportResponse{Info: "push data failed"}, nil
			fmt.Println(err.Error())
			return
		}
	} else {
		err := RedisClient().HSet(in.Key, string(in.Timestamp), in.Loss).Err()
		if err != nil {
			logger.Error("push data failed", zap.String("detail", err.Error()),
				zap.String("key", in.Key))
			//return &pb.ReportResponse{Info: "push data failed"}, nil
			fmt.Println(err.Error())
			return
		}
	}

	d, err := RedisClient().TTL(in.Key).Result()
	if err != nil {
		//return &pb.ReportResponse{Info: "push failed"}, err
		fmt.Println(err.Error())
		return
	}
	if d <= 0 {
		err := RedisClient().Expire(in.Key, time.Duration(in.ExpireTime) * time.Second).Err()
		if err != nil {
			logger.Error("set expire failed", zap.String("detail", err.Error()),
				zap.String("key", in.Key))
			fmt.Println(err.Error())
			return
		}
	}

	//fmt.Printf("push to redis success!")
	pool.Done()
	return
}
