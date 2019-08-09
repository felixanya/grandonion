package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"
)

// export
//var TaskPlugin *TimerTask
// TODO 这里使用指针会导致symbol绑定到Plugin失败
var TaskPlugin TimerTask

type Input struct {
	Interval    int     `json:"interval"`
	Content     string  `json:"content"`
}

type Output struct {
	ExecTimes       int     `json:"exec_times"`
	AgentName       string  `json:"agent_name"`
	Info            string  `json:"info"`
}

type TimerTask struct {
	available   bool

	done        chan bool
}

func (t *TimerTask) Load() error {
	TaskPlugin = TimerTask{
		available: false,
		done: make(chan bool),
	}
	return nil
}

func (t *TimerTask) Available() bool {
	return t.available
}

func (t *TimerTask) Enable() error {
	t.available = true

	return nil
}

func (t *TimerTask) Disable() error {
	t.available = false

	t.done <- true
	return nil
}

func (t *TimerTask) Execute(ctx context.Context, input ...[]byte) ([]byte, error) {
	if !t.available {
		return nil, errors.New("plugin is not available")
	}

	times := 0
	interruptInfo := ""

	if len(input) == 0 {
		return nil, errors.New("缺少参数")
	} else if len(input) > 1 {
		return nil, errors.New("too many parameters")
	}

	inputParm := &Input{}
	if err := json.Unmarshal(input[0], inputParm); err != nil {
		return nil, err
	}

OUTFOR:
	for {
		select {
		case <-time.After(time.Duration(inputParm.Interval) * time.Second):
			// TODO 任务超时控制
			log.Println(">>>执行任务>>>", inputParm.Content)
		case <-t.done:
			interruptInfo = "interrupted by manual close"
			break OUTFOR
		case <-ctx.Done():
			interruptInfo = "interrupted by context"
			break OUTFOR
		}
		times += 1
	}

	output := struct {
		Status      string  `json:"status"`
		Message     string  `json:"message"`
		Data        Output  `json:"data"`
	}{
		Status: "success",
		Message: "Execute task success",
		Data: Output{
			ExecTimes: times,
			AgentName: "test_agent",
			Info: interruptInfo,
		},
	}

	bt, err := json.Marshal(&output)
	if err != nil {
		return nil, err
	}

	return bt, nil
}

