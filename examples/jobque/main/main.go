package main

import (
	"fmt"
	"github.com/aaronize/grandonion/examples/jobque"
	"time"
)

var workerPool chan chan jobque.Job

func main() {
	// 注意 这里记得workerPool是要带缓存的chan，不然会死锁
	workerPool = make(chan chan jobque.Job, 100)
	worker := jobque.NewWorker(workerPool)

	worker.Start()

	for i := 0; i < 10; i++ {
		//payload := Payload{}
		worker.JobChannel <- jobque.Job{}

		time.Sleep(1 * time.Second)
	}
	fmt.Println("finish. shutdown worker")
	worker.Stop()
}
