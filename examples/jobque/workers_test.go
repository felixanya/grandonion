package jobque

import (
	"fmt"
	"testing"
	"time"
)

func TestNewWorker(t *testing.T) {
	var workerPool chan chan Job
	workerPool = make(chan chan Job, 100)

	worker := NewWorker(workerPool)

	worker.Start()

	for i := 0; i < 10; i++ {
		//payload := Payload{}
		worker.JobChannel <- Job{}

		time.Sleep(2 * time.Second)
	}

	fmt.Println("execute finish>>>")

	worker.Stop()
}