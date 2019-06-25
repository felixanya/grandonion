package main

import (
	"fmt"
	"time"
)

func worker2(ch chan string) {
	go func() {
		for {
			t := time.NewTimer(3 * time.Second)

			select {
			case n := <- ch:
				fmt.Println(">>>", n)
				time.Sleep(1 * time.Second)
			case <-t.C:
				fmt.Println("timeout.")
			}
		}
	}()

	return
}

func main() {
	// 从一个已经关闭的channel中取值，会一直读取到零值
	ch := make(chan string)

	worker2(ch)
	time.Sleep(1 * time.Second)
	ch <- "30"
	time.Sleep(10 * time.Second)
	close(ch)

	time.Sleep(30 * time.Second)
}
