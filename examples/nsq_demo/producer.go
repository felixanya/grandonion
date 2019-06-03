package nsq_demo

import (
	"fmt"
	"github.com/nsqio/go-nsq"
)

//var producer *nsq.Producer

type ProducerT struct{
	address 	string
	producer *nsq.Producer
}

func (p *ProducerT) InitProducer() error {
	fmt.Println("address:", p.address)
	var err error
	p.producer, err = nsq.NewProducer(p.address, nsq.NewConfig())
	if err != nil {
		panic(err)
	}
	fmt.Printf("connect %s success.\n", p.address)
	return nil
}

func (p *ProducerT) Publish(topic, message string) error {
	var err error
	if p.producer != nil {
		if message == "" {
			return fmt.Errorf("message is empty")
			//message = "hellow"
		}
		err = p.producer.Publish(topic, []byte(message))
		return err
	}

	return fmt.Errorf("producer is nil")
}

func (p *ProducerT) Stop() {
	p.producer.Stop()
}
