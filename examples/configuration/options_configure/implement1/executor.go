package implement1

import (
	"context"
	"fmt"
	"sync"
)

type Executor struct {
	executorId 		string
	ctx 			context.Context
	cancel 			context.CancelFunc

	eOptions executorOptions
	eLock    sync.RWMutex
}

func (e *Executor) PrintExecutorInfo() {
	fmt.Printf(">>>%+v", e)
}

func (e *Executor) PrintExecutorOptions() {
	fmt.Printf(">>>Options>>>%+v", e.eOptions)
}
