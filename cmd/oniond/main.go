package main

import (
	"fmt"
	"github.com/judwhite/go-svc/svc"
	"sync"
	"syscall"
)

type program struct {
	once 	sync.Once
}

func (p *program) Init(env svc.Environment) error {
	if env.IsWindowsService() {
		fmt.Println("windows environment")
	}
	return nil
}

func (p *program) Start() error {

	return nil
}

func (p *program) Stop() error {

	return nil
}

func main() {
	p := &program{}
	if err := svc.Run(p, syscall.SIGINT, syscall.SIGKILL); err != nil {
		fmt.Println("error occur,", err.Error())
		return
	}
}

