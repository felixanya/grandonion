package main

import (
	"fmt"
	"time"
)

func main3() {
	ch := make(chan int)

	ch <- 1

	close(ch)

	go func() {
		for n := range ch {
			fmt.Println(n)
		}
	}()

	time.Sleep(10 * time.Second)
}
