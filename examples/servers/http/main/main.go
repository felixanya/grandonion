package main

import (
	"fmt"
	http2 "github.com/aaronize/grandonion/examples/servers/http"
)

func main() {
	confPath := ""
	server := http2.NewServer(confPath)
	if err := server.RunServer(); err != nil {
		fmt.Println("start server error", err.Error())
		return
	}
}

