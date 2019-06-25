package main

import (
	"fmt"
	"time"
)

var VMTaskQue chan map[string][]string

func NewTaskQue(size int) chan map[string][]string {
	ch := make(chan map[string][]string, size)
	return ch
}

func main() {
	VMTaskQue = NewTaskQue(10)
	var done = make(chan bool)

	go func() {
		for {
			select {
			case task := <- VMTaskQue:
				fmt.Println(">>>receive task>>>", task)
			case <- time.After(3 * time.Second):
				fmt.Println(">>>timeout>>>")
			case <- done:
				fmt.Println(">>>return>>>")
				return
			}
		}
	}()

	t1 := make(map[string][]string)
	t1["create"] = []string{"ni", "hao", "shi", "jie"}
	VMTaskQue <- t1

	time.Sleep(10 * time.Second)
	done <- true
	time.Sleep(5 * time.Second)
}
