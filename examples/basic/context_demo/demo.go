package main

import (
	"context"
	"fmt"
	"time"
)

//type CustomContext struct {
//
//}
//
//func (cc *CustomContext) Deadline() (deadline time.Time, ok bool)  {
//	return time.Now(), false
//}
//
//func (cc *CustomContext) Done() <-chan struct{} {
//	return make(<-chan struct{})
//}
//
//func (cc *CustomContext) Err() error {
//	return nil
//}
//
//func (cc *CustomContext) Value(key interface{}) interface{} {
//	return nil
//}

func main() {
	env := map[string]string{"GOOS": "linux", "GOARCH": "amd64", "CGO_ENABLE": "0"}
	// 装饰器模式？建造者模式？
	ctx, cancel := context.WithCancel(context.WithValue(context.Background(), "env", env))
	for i := 0; i < 5; i++ {
		go func(ctx context.Context) {
			for {
				select {
				case <- ctx.Done():
					fmt.Println("stop...")
					return
				default:
					fmt.Println("goroutine running...GOOS:", ctx.Value("env").(map[string]string)["GOOS"])
					time.Sleep(2 * time.Second)
				}
			}
		}(ctx)
	}

	time.Sleep(10 * time.Second)
	fmt.Println("notify to stop")
	cancel()

	time.Sleep(5 * time.Second)
}
