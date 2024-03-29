package nsq_demo

import (
	"fmt"
	"testing"
	"time"
)

func TestProducerT_InitProducer(t *testing.T) {
	pt := &ProducerT{
		// nsqd
		address: "127.0.0.1:4150",
	}
	InitProducer()

	for i := 0; i < 10; i++ {
		Publish("topic4test", "testing messages")
		fmt.Printf(">>>>>>published %d>>>\n", i)
		time.Sleep(5 * time.Second)
	}

	Stop()
}
