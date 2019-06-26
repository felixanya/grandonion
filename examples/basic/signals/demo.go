package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ch := make(chan os.Signal)
	//signal.Notify(ch) // 接收所有信号
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		for s := range ch {
			switch s {
			case syscall.SIGTERM:
				log.Println("terminate:", s)
				os.Exit(0)
			case syscall.SIGHUP:
				log.Println("sighup:", s)
				os.Exit(0)
			case syscall.SIGINT:
				log.Println("sigint: ", s)
				os.Exit(0)
			default:
				log.Println("other:", s)
				os.Exit(0)
			}
		}
	}()

	log.Println("进程启动...")
	time.Sleep(30 * time.Second)
}
