package up2redis

import (
	"github.com/aaronize/grandonion/examples/rpc/grpc_demo/up2redis/toredis"
	"testing"
)

func TestRegisterRedisServer(t *testing.T) {
	s := &Server{}
	s.msgChan = make(chan *redis.ReportRequest, 10)



}
