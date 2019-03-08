package jobque

import (
	"fmt"
	"os"
)

var(
	MaxWorker = os.Getenv("MAX_WORKERS")
	MaxQueue = os.Getenv("MAX_QUEUE")
)

type Worker struct {
	WorkerPool 	chan chan Job
	JobChannel	chan Job
	quit 		chan bool
}

func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit: make(chan bool),
	}
}

func (w Worker) Start() {
	go func() {

		for {
			// register the current worker into the worker queue.
			w.WorkerPool <- w.JobChannel
			fmt.Println(">>>add jobChannel to workerPool>>>success.")

			select {
			case job := <- w.JobChannel:
				// receive a process request
				fmt.Println("处理请求", job.doJob())
			case <- w.quit:
				// receive a stop signal
				return
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
