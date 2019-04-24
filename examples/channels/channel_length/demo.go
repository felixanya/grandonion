package main

import "fmt"

func main() {
	ch := make(chan struct{}, 5)
	ch <- struct{}{}

	fmt.Println(len(ch))
}
