package main

import (
	"bufio"
	"fmt"
	"github.com/Shopify/sarama"
	"io"
	"os"
	"time"
)

var (
	FilePath = ""
	SendChanSize = 100000
	ProducerNum = 10
	Topic = "test_pm"
	Hosts = []string{""}
)

type KafkaClient struct {
	addr 		[]string
	num 		int
	config 		*sarama.Config
	msgChan 	chan *sendMsg

	quit 		chan bool
}

type sendMsg struct {
	topic 	string
	value 	[]byte
}


func main() {
	kafkaClient := NewKafkaClient(Hosts)
	if err := kafkaClient.InitAsyncProducers(); err != nil {
		panic(err)
	}

	var upload2Que = func(sendMsg *sendMsg) {
		//
		kafkaClient.msgChan <- sendMsg
	}

	if err := ReadFile(FilePath, upload2Que); err != nil {
		panic(err)
	}

	time.Sleep(5 * time.Second)
}

func ReadFile(filePath string, handle func(*sendMsg)) error {
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)

	for {
		line, _, err := buf.ReadLine()

		handle(&sendMsg{
			topic: Topic,
			value: line,
		})

		if err != nil {
			if err == io.EOF{
				return nil
			}
			return err
		}
		return nil
	}
}

func NewKafkaClient(addr []string) *KafkaClient {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Compression = sarama.CompressionSnappy
	config.Version = sarama.V2_1_0_0

	return &KafkaClient{
		addr: addr,
		config: config,
		num: ProducerNum,
		msgChan: make(chan *sendMsg, SendChanSize),
		quit: make(chan bool),
	}
}

func (kc *KafkaClient) InitAsyncProducers() error {
	if kc.num == 0 {
		return kc.GenerateSingleAsyncProducer()
	}

	for i := 0; i < kc.num; i++ {
		kc.GenerateSingleAsyncProducer()
	}

	return nil
}

func (kc *KafkaClient) GenerateSingleAsyncProducer() error {
	producer, err := sarama.NewAsyncProducer(kc.addr, kc.config)
	if err != nil {
		return err
	}

	go func(p sarama.AsyncProducer) {
		for {
			select {
			case err := <- p.Errors():
				if err != nil {
					fmt.Printf("[ERROR] %s input queue failed, %s\n", time.Now().Format("2006-01-02 15:04:05.999999999"), err.Error())
				}
			case <- p.Successes():
				fmt.Printf("[INFO] %s input queue success. \n", time.Now().Format("2006-01-02 15:04:05.999999999"))
			case <- kc.quit:
				fmt.Println("quit result listener")
				return
			}
		}
	}(producer)

	go func(p sarama.AsyncProducer) {
		for {
			select {
			case msg := <- kc.msgChan:
				toSend := &sarama.ProducerMessage{
					Topic: msg.topic,
					Value: sarama.ByteEncoder(msg.value),
					Timestamp: time.Now(),
				}
				p.Input() <- toSend
			case <- kc.quit:
				fmt.Println("quit producer")
				return
			}
		}
	}(producer)

	return nil
}

func (kc *KafkaClient) CloseProducer() {
	close(kc.quit)
}
