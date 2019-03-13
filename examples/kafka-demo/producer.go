package kafka_demo

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
)

type ProducerPool struct {
	InputQueue 		chan []byte
	ProducerNum 	int

	quit 			chan bool
}

func (pp *ProducerPool) Run() {

	for i := 0; i < pp.ProducerNum; i++ {
		go pp.genAsyncProducer()
	}
}

func (pp *ProducerPool) genAsyncProducer() error {
	// TODO 配置
	producer, err := sarama.NewAsyncProducer([]string{""}, sarama.NewConfig())
	if err != nil {
		return err
	}
	defer producer.AsyncClose()

	go func(p sarama.AsyncProducer) {
		for {
			select {
			case err := <- p.Errors():
				if err != nil {
					fmt.Println("")
				}
			case <- p.Successes():
				fmt.Println("推送kafka成功")

			// TODO 退出信号
			}
		}
	}(producer)

	for {
		select {
		case pb := <- pp.InputQueue:
			plainData := make(map[string]interface{})
			if err := json.Unmarshal(pb, &plainData); err != nil {
				fmt.Println("json parse error", err.Error())
				continue
			}

			var topic string= ""

			msg := &sarama.ProducerMessage{
				Topic: topic,
				Value: sarama.ByteEncoder(pb),
			}

			producer.Input() <- msg

		}
	}
}
