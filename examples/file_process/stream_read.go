package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/Shopify/sarama"
	"io"
	"io/ioutil"
	"os"
	"time"
)

var FilePath string

type Config struct {
	FilePath 		string		`json:"FilePath"`
	SendChanSize 	int			`json:"SendChanSize"`
	ProducerNum  	int			`json:"ProducerNum"`
	Topic 			string		`json:"Topic"`
	Hosts 			[]string	`json:"Hosts"`
	WaitTime 		int64 		`json:"WaitTime"`
}

type KafkaClient struct {
	addr 		[]string
	num 		int
	config 		*sarama.Config
	msgChan 	chan *sendMsg

	quit 		chan bool
	done 		chan bool
}

type sendMsg struct {
	topic 	string
	value 	[]byte
}

var conf *Config

func main() {
	flag.Parse()
	if FilePath == "" {
		panic("需要指定配置文件路径: ./upload -c /config/file/path")
	}

	conf = &Config{}
	if err := conf.LoadConfiguration(FilePath); err != nil {
		panic(err)
	}

	kafkaClient := NewKafkaClient(conf.Hosts)
	if err := kafkaClient.InitAsyncProducers(); err != nil {
		panic(err)
	}

	var upload2Que = func(sendMsg *sendMsg) {
		//
		kafkaClient.msgChan <- sendMsg
	}

	if err := ReadFile(conf.FilePath, upload2Que); err != nil {
		panic(err)
	}

	<- kafkaClient.done
	//time.Sleep(time.Duration(conf.WaitTime) * time.Second)
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
		if err != nil {
			if err == io.EOF {
				if err := os.Truncate(filePath, 0); err != nil {
					return err
				}
				return nil
			}
			return err
		}

		handle(&sendMsg{
			topic: conf.Topic,
			value: line,
		})
	}

	return nil
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
		num: conf.ProducerNum,
		msgChan: make(chan *sendMsg, conf.SendChanSize),
		quit: make(chan bool),
		done: make(chan bool),
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
		timer := time.NewTimer(time.Duration(conf.WaitTime) * time.Second)
		defer timer.Stop()

		for {
			select {
			case err := <- p.Errors():
				if err != nil {
					fmt.Printf("[ERROR] %s input queue failed, %s\n", time.Now().Format("2006-01-02 15:04:05.999999999"), err.Error())
				}
				timer.Reset(time.Duration(conf.WaitTime) * time.Second)
			case <- p.Successes():
				fmt.Printf("[INFO] %s input queue success. \n", time.Now().Format("2006-01-02 15:04:05.999999999"))
				timer.Reset(time.Duration(conf.WaitTime) * time.Second)
			case <- kc.quit:
				fmt.Println("quit result listener")
				return
			case <- timer.C:
				kc.done <- true
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

func (kc *KafkaClient) CloseProducers() {
	close(kc.quit)
}

func (c *Config) LoadConfiguration(path string) error {
	if path == "" {
		path = "./config.json"
	}
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	if len(content) == 0 {
		return errors.New("config file is empty. ")
	}

	if err := json.Unmarshal(content, c); err != nil {
		return err
	}

	return nil
}

func init() {
	flag.StringVar(&FilePath, "c", "./config.json", "指定配置文件路径")
}
