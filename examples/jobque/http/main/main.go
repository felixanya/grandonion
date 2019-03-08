package main

import (
	"fmt"
	"github.com/aaronize/grandonion/examples/jobque/http"
)

func main() {
	confPath := ""
	server := http.NewServer(confPath)
	if err := server.RunServer(); err != nil {
		fmt.Println("start server error", err.Error())
		return
	}
}

