package main

import (
	"fmt"
	"github.com/RussellLuo/timingwheel"
	"time"
)

func main() {
	tw := timingwheel.NewTimingWheel(time.Millisecond, 20)
	tw.Start()
	defer tw.Stop()

	exitC := make(chan time.Time, 1)
	tw.AfterFunc(time.Second * 10, func() {
		fmt.Println("The timer fires")
		exitC <- time.Now()
	})

	<-exitC
}
