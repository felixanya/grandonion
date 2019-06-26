package main

import (
	"fmt"
	"github.com/aaronize/grandonion/examples/basic/condition_compile/json_test"
	"github.com/aaronize/grandonion/examples/basic/condition_compile/server"
	"log"
)

type User struct {
	Name 	string
	Age 	int
}

func main() {
	u := &User{"aaron", 20}
	b, err := json_test.MarshalIdent(u, "", "")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(string(b))
	}

	s := server.Server{}
	if err := s.Start(); err != nil {
		log.Fatalf(err.Error())
	}
}
