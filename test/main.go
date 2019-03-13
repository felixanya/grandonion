package main

import (
	"fmt"
	"time"
)

var JobQueue chan *Job

type Job struct{
	State 	string
}

func main() {
	JobQueue = make(chan *Job, 10)

	go handler(1)
	j := Job{State: "creating"}
	JobQueue <- &j

	fmt.Printf("print in main func: %p \n", &j)
	select {}
}

func handler(timeout time.Duration) {

	select {
	case job := <- JobQueue:
		fmt.Printf("print in handler: %p \n", job)
		return
	case <- time.After(timeout * time.Second):
		fmt.Println("超时了")
		return
	}
}

