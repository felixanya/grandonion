package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/Shopify/sarama"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var (
	ConfFilePath string
	FilePath 	 string

	successCount int64
	failCount    int64

	conf         *Config
	hasFail      bool
	failedChan   chan []byte
)

type Config struct {
	SendChanSize 	int			`json:"SendChanSize"`
	FailedChaneSize int         `json:"FailedChaneSize"`
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
	//done 		chan bool
	wg          *sync.WaitGroup
}

type sendMsg struct {
	topic 	string
	value 	[]byte
}

func main() {
	defer func() {
		fmt.Printf("{success:%d, fail:%d}\n", atomic.LoadInt64(&successCount), atomic.LoadInt64(&failCount))
	}()

	flag.Parse()
	if ConfFilePath == "" {
		panic("需要指定配置文件路径: ./upload -c /config/file/path")
	}

	conf = &Config{}
	if err := conf.LoadConfiguration(ConfFilePath); err != nil {
		panic(err)
	}

	failedChan = make(chan []byte, conf.FailedChaneSize)
	hasFail = false

	kafkaClient := NewKafkaClient(conf.Hosts)
	if err := kafkaClient.InitAsyncProducers(); err != nil {
		panic(err)
	}

	var upload2Que = func(sendMsg *sendMsg) {
		//
		kafkaClient.msgChan <- sendMsg
	}

	if err := ReadFileUseScanner(FilePath, upload2Que); err != nil {
		panic(err)
	}

	//<- kafkaClient.done
	kafkaClient.wg.Wait()

	//time.Sleep(time.Duration(conf.WaitTime) * time.Second)
	kafkaClient.CloseProducers()

	if !hasFail {
		return
	}

	f, err := os.OpenFile(FilePath, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("写入文件失败, %s", err.Error())
	}
	defer f.Close()

	buf := bufio.NewWriter(f)

	//for {
	//	if bt, ok := <- failedChan; ok {
	//		if _, err := buf.WriteString(string(bt)); err != nil {
	//			log.Println("写入文件失败,", err.Error())
	//		}
	//	} else {
	//		break
	//	}
	//}

	//ticker := time.NewTicker(2 * time.Second)
	//defer ticker.Stop()
	timer := time.NewTimer(1 * time.Second)

OUTTER:
	for {
		select {
		case bt := <-failedChan:
			if _, err := buf.WriteString(string(bt)); err != nil {
				log.Println("写入文件失败,", err.Error())
			}
			timer.Reset(1 * time.Second)
		case <-timer.C:
			break OUTTER
		}
	}

	if err := buf.Flush(); err != nil {
		log.Println("写入文件失败,", err.Error())
	}

	return
}

func ReadFileUseScanner(filePath string, handle func(*sendMsg)) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		line = strings.Trim(line, "\n")

		//var data map[string]interface{}
		//if err := json.Unmarshal([]byte(line), &data); err != nil {
		//	continue
		//}

		handle(&sendMsg{
			topic: conf.Topic,
			value: []byte(line),
		})
	}
	//if err := os.Truncate(filePath, 0); err != nil {
	//	return err
	//}
	if err := os.Remove(filePath); err != nil {
		log.Println("remove file error, ", err.Error())
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
		wg: &sync.WaitGroup{},
	}
}

func (kc *KafkaClient) InitAsyncProducers() error {
	if kc.num == 0 {
		return kc.GenerateSingleAsyncProducer()
	}

	for i := 0; i < kc.num; i++ {
		err := kc.GenerateSingleAsyncProducer()
		if err != nil {
			log.Fatal(err)
		}
		kc.wg.Add(1)
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
				atomic.AddInt64(&failCount, 1)
				if err != nil {
					fmt.Printf("[ERROR] %s input queue failed, %s\n", time.Now().Format("2006-01-02 15:04:05.999999999"), err.Error())
				}
				bt, _ := err.Msg.Value.Encode()
				failedChan <- bt
				hasFail = true

				timer.Reset(time.Duration(conf.WaitTime) * time.Second)
			case <- p.Successes():
				atomic.AddInt64(&successCount, 1)
				//fmt.Printf("[INFO] %s input queue success. \n", time.Now().Format("2006-01-02 15:04:05.999999999"))
				timer.Reset(time.Duration(conf.WaitTime) * time.Second)
			case <- kc.quit:
				//fmt.Println("quit result listener")
				//p.AsyncClose()
				return
			case <- timer.C:
				//kc.done <- true
				kc.wg.Done()
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
				//fmt.Println("quit producer")
				p.AsyncClose()
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
	flag.StringVar(&ConfFilePath, "c", "./producer_config.json", "指定配置文件路径")
	flag.StringVar(&FilePath, "f", "", "指定上传至kafka的数据文件路径")
}
