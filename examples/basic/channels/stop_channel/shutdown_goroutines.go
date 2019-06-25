package main

import (
	"fmt"
	"time"
)

var done chan int

func main() {
	done = make(chan int)
	fmt.Println("before start goroutine...")
	startGoroutine(10)
	fmt.Println("after start goroutine...")
	time.Sleep(3 * time.Second)

	// 关闭channel
	close(done)

	time.Sleep(20 * time.Second)
}

func startGoroutine(num int) {
	for i := 0; i < num; i++ {
		go func() {
			select {
			case <- done:
				fmt.Println("收到终止信号，退出goroutine...")
				return
			}
		}()
	}
}

