package main

import (
	"fmt"
	"time"
)

func handleTask() <-chan string {
	ch := make(chan string)
	go func() {
		time.Sleep(6 * time.Second)
		ch <- "mission completed"
	}()

	return ch
}

func main() {
	//for {
	select {
	case msg := <-handleTask():
		fmt.Println(msg)
	case <-time.After(5 * time.Second):
		fmt.Print("Timeout")
		return
	}
	//}
}
