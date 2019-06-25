package nsq_demo

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"sync"
	"time"
)

type ConsumerT struct {
	topic 		string
	channel 	string
	address 	string

	sig 		chan bool
}

//func NsqConsumeMsg() error {
//
//	return nil
//}

func (c ConsumerT) HandleMessage(msg *nsq.Message) error {
	fmt.Println("receive", msg.NSQDAddress, "message:", string(msg.Body))
	return nil
}

func (c *ConsumerT) NsqConsumers() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = time.Second
	consumer, err := nsq.NewConsumer(c.topic, c.channel, cfg)
	if err != nil {
		panic(err)
	}
	consumer.SetLogger(nil, 0)
	consumer.AddHandler(&ConsumerT{})

	if err := consumer.ConnectToNSQD(c.address); err != nil {
		panic(err)
	}

	<- c.sig
}
