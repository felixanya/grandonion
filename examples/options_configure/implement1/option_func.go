package implement1

import (
	"context"
	"github.com/satori/go.uuid"
	"sync"
	"time"
)

type Identifier func(string) error

// 利用闭包实现配置
// 这种方法既很好的封装了配置字段和结构体，外部只能通过提供的函数来对指定的配置信息进行修改。
// 这种方式在grpc中的Dial的配置连接信息中有所应用，
// 另外在httputil的配置连接信息中也有十分类似的实现。
// 对比发现，以下实现比httputil中的实现封装的更好。
type executorOptions struct {
	action 			string
	isEmergency 	bool
	timeout 		time.Duration
	identifier 		Identifier
}

// 执行器配置接口
type ExecutorOption interface {
	apply(*executorOptions)
}

// empty option
type EmptyExecutorOption struct {}

func (eeo *EmptyExecutorOption) apply(eo *executorOptions) {
}

//
type funcExecutorOption struct {
	fn 		func(*executorOptions)
}

func (feo *funcExecutorOption) apply(eo *executorOptions) {
	feo.fn(eo)
}

func newFuncExecutorOption(fn func(*executorOptions)) *funcExecutorOption {
	return &funcExecutorOption{
		fn: fn,
	}
}

// 使用函数配置开放的配置信息
func WithEmergency() ExecutorOption {
	return newFuncExecutorOption(func(ops *executorOptions) {
		ops.isEmergency = true
	})
}

func WithAction(a string) ExecutorOption {
	return newFuncExecutorOption(func(ops *executorOptions) {
		ops.action = a
	})
}

func WithTimeout(t time.Duration) ExecutorOption {
	return newFuncExecutorOption(func(ops *executorOptions) {
		ops.timeout = t
	})
}

//
func Execute(opts ...ExecutorOption) (executor *Executor, err error) {
	return ExecuteContext(context.Background(), opts...)
}
//
func ExecuteContext(ctx context.Context, opts ...ExecutorOption) (executor *Executor, err error) {
	uid, err := uuid.NewV4()

	// 默认配置
	executor = &Executor{
		executorId: uid.String(),
		ctx: ctx,
		eLock: sync.RWMutex{},
	}

	for _, opt := range opts {
		opt.apply(&executor.eOptions)
	}

	return executor, nil
}
