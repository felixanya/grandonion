package up2redis

import (
	"context"
	pb "github.com/aaronize/grandonion/examples/grpc_demo/up2redis/toredis"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

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
	pb.RegisterRedisServer(rs, &Server{})
	if err := rs.Serve(li); err != nil {
		log.Fatalf("Failed to serve, %s", err.Error())
	}
	return nil
}

func (s *Server) Report(ctx context.Context, in *pb.ReportRequest) (*pb.ReportResponse, error) {
	// upload to redis
	if in.Delay == 0.0 {
		err := RedisClient().HSet(in.Key, string(in.Timestamp), in.Delay).Err()
		if err != nil {
			logger.Error("push data failed", zap.String("detail", err.Error()),
				zap.String("key", in.Key))
			return &pb.ReportResponse{Info: "push data failed"}, nil
		}
	} else {
		err := RedisClient().HSet(in.Key, string(in.Timestamp), in.Loss).Err()
		if err != nil {
			logger.Error("push data failed", zap.String("detail", err.Error()),
				zap.String("key", in.Key))
			return &pb.ReportResponse{Info: "push data failed"}, nil
		}
	}

	d, err := RedisClient().TTL(in.Key).Result()
	//log.Printf(">>>>>>>>>>>>>expire time：%d>>>>>>>", d)
	if err != nil {
		return &pb.ReportResponse{Info: "push failed"}, err
	}
	if d <= 0 {
		err := RedisClient().Expire(in.Key, time.Duration(in.ExpireTime) * time.Second).Err()
		if err != nil {
			logger.Error("set expire failed", zap.String("detail", err.Error()),
				zap.String("key", in.Key))
			return &pb.ReportResponse{Info: "set expire failed"}, nil
		}
	}

	return &pb.ReportResponse{Info: "push data success"}, nil
}
