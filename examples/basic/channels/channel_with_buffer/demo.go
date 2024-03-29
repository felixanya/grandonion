package main

import (
	"fmt"
	"time"
)

/*
 带缓冲的channel
	chan为空时，会阻塞读
	chan为满时，会阻塞写
 */
func main() {
	// 带缓冲的channel
	ch := make(chan string, 2)
	n := 5

	go func() {
		time.Sleep(5 * time.Second)
		for {
			select {
			case s := <- ch:
				fmt.Println("get string from channel: ", s)
				time.Sleep(2 * time.Second)
			case <- time.After(20 * time.Second):
				fmt.Println(">>>timeout>>>")
				return
			}
		}
	}()

	for i := 0; i < n; i++ {
		ch <- fmt.Sprintf("第%d个", i)
		fmt.Printf(">>>>>>>>>%d>>>>>>>>>\n", i)
	}

	time.Sleep(30 * time.Second)
	fmt.Println(">>>FINISH>>>")
}
