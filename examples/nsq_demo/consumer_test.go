package nsq_demo

import (
	"testing"
)

func TestNsqConsumeMsg(t *testing.T) {
	consumer := &ConsumerT{
		topic: "topic4test",
		channel: "chan4test",
		// nsqd
		address: "127.0.0.1:4150",
	}

	consumer.NsqConsumers()
}
