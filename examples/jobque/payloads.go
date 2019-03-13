package jobque

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
	Queue chan Payload
	MAX_QUEUE = 1000
	MaxLength int64 = 5000
)

type Payload struct {
	AllocRes 	func() error
}

type PayloadCollection struct {
	Version 	string 	`json:"version"`
	Token 		string	`json:"token"`
	Payloads 	[]Payload `json:"data"`
}

func init() {
	Queue = make(chan Payload, MAX_QUEUE)
}

func (p *Payload) UploadToS3() error {
	fmt.Println("execute upload")

	return nil
}

func payloadHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		res.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("Handling payloads...")
	var content = &PayloadCollection{}

	err := json.NewDecoder(io.LimitReader(req.Body, MaxLength)).Decode(content)
	if err != nil {
		res.Header().Set("Content-Type", "application/json; charset=UTF-8")
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, payload := range content.Payloads {
		// 直接使用goroutine
		// go payload.UploadToS3()

		//
		//Queue <- payload

		// 创建一个job
		worker := Job{Payload: payload}
		// 推JobQueue队列中
		JobQueue <- worker
	}

	res.WriteHeader(http.StatusOK)
}

func StartProcessor() {
	for {
		select {
			case job := <- Queue:
				job.UploadToS3()
		}
	}
}
