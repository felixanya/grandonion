package up2redis

import "testing"
import pb "github.com/aaronize/grandonion/examples/grpc_demo/up2redis/toredis"

func TestRegisterRedisServer(t *testing.T) {
	s := &Server{}
	s.msgChan = make(chan *pb.ReportRequest, 10)



}
