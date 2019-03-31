package main

import (
	"flag"
	pb "github.com/aaronize/grandonion/examples/grpc_demo/route_guide/routeguide"
)

var (
	tls = flag.Bool("tls", false, "Connection uses TLS if true, else plain tcp")
	caFile = flag.String("ca_file", "", "The file containing  the CA root cert file")
	serverAddr = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name to verify the hostname return by TLS handshake")
)

func printFeature(client pb.RouteGuideClient)  {
	
}
