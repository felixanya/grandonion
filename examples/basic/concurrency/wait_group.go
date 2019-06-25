package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("No.1")
		wg.Done()
	}()

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("No.2")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("done")
}
