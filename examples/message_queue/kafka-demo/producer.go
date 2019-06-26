package kafka_demo

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
)

type ProducerPool struct {
	InputQueue 		chan []byte
	ProducerNum 	int

	quit 			chan bool
}

func (pp *ProducerPool) Run() {

	//for i := 0; i < pp.ProducerNum; i++ {
	//	go pp.genAsyncProducer()
	//}
	if err := pp.genAsyncProducer(); err != nil {
		log.Fatal(err)
	}
}

func (pp *ProducerPool) genAsyncProducer() error {
	// TODO 配置
	producer, err := sarama.NewAsyncProducer([]string{""}, sarama.NewConfig())
	if err != nil {
		return err
	}
	//defer producer.AsyncClose()
	//fmt.Printf("producer out: %p", producer)
	log.Printf("producer out: %p", producer)

	// 监听结果
	go func(p sarama.AsyncProducer) {
		//fmt.Printf("producer in: %p", p)
		log.Printf("producer in: %p", p)
		for {
			select {
			case err := <- p.Errors():
				if err != nil {
					fmt.Println(err)
				}
			case <- p.Successes():
				fmt.Println("推送kafka成功")
			case <- pp.quit:
				// TODO IMPORTANT！！不要在这里关闭producer，要在input消息的goroutine里关闭！
				//p.AsyncClose()
				return
			}
		}
	}(producer)

	go func(p sarama.AsyncProducer) {
		for {
			select {
			case pb := <- pp.InputQueue:
				plainData := make(map[string]interface{})
				if err := json.Unmarshal(pb, &plainData); err != nil {
					fmt.Println("json parse error", err.Error())
					continue
				}
				var topic = "testing"
				msg := &sarama.ProducerMessage{
					Topic: topic,
					Value: sarama.ByteEncoder(pb),
				}

				producer.Input() <- msg
			case <- pp.quit:
				// TODO 必须要在这里关闭producer！！
				// 调用AsyncClose()函数时，会使用close()函数显式关闭producer中的channel，如success, errors, input, retry等。
				// channel的关闭使这里监听input管道的goroutine短时间内收到一些零值，上报到队列。
				p.AsyncClose()
				return
			}
		}
	}(producer)

	return nil
}
