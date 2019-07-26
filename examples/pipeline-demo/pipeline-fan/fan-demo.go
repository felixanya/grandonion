package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, num := range nums {
			out <- num
		}
	}()

	return out
}

func square(inChan <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range inChan {
			time.Sleep(1 * time.Second)
			out <- num
		}
	}()

	return out
}

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup

	collect := func(in <-chan int) {
		defer wg.Done()
		for num := range in {
			out <- num * num
		}
	}

	wg.Add(len(cs))

	// FAN-IN
	for _, c := range cs {
		go collect(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	in := producer(1, 2, 3, 4, 5)

	// FAN-OUT
	c1 := square(in)
	c2 := square(in)
	c3 := square(in)

	for ret := range merge(c1, c2, c3) {
		fmt.Printf("%3d \n", ret)
	}
	fmt.Println()
}
