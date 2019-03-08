package main

import "fmt"

var JobQPool chan chan string

type JobQ struct {
	job chan string
}

func main() {
	job1 := JobQ{job: make(chan string, 10)}
	job2 := JobQ{job: make(chan string, 10)}

	JobQPool = make(chan chan string, 10)

	job1.job <- "job1 string"
	JobQPool <- job1.job

	fmt.Println("jobque:", <-<-JobQPool)

	var result string
	job2.job <- "job2 string"
	result = <- job2.job

	fmt.Println("result:", result)

	var notInitChan = make(chan string)
	go func() {
		result2 := <-notInitChan
		fmt.Println("this is:", result2)
	}()
	notInitChan <- "not init channel"

}
