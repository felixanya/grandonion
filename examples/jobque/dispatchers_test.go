package jobque

import "testing"

func TestNewDispatcher(t *testing.T) {
	JobQueue = make(chan Job)

	dispatcher := NewDispatcher(100)

	dispatcher.Run()

	for i := 0; i < 10; i++ {
		job := Job{}
		JobQueue <- job
	}
}
