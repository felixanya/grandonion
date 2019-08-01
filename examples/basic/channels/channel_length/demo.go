package main

import "fmt"

func main() {
	ch := make(chan struct{}, 5)
	ch <- struct{}{}

	// channel length
	fmt.Println(len(ch))
	// channel capacity
	fmt.Println(cap(ch))
}
