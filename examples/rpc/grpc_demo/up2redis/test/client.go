package main

import (
	"context"
	"github.com/aaronize/grandonion/examples/rpc/grpc_demo/up2redis/toredis"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:8712"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connect rpc server error. %s", err.Error())
	}
	defer conn.Close()

	c := redis.NewRedisClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	report := &redis.ReportRequest{
		Delay: 10,
		Timestamp: 1559044501,
		Loss: 20,
		Key: "420e57b017066b44e05ea1577f6e2e12",
		ExpireTime: 60,
	}

	r, err := c.Report(ctx, report)
	if err != nil {
		log.Fatalf("Report error. %s", err.Error())
	}
	log.Printf("Report success! Response: %s", r.Info)
}
