package main

import (
    "log"
    "time"
)
/*
for range channel 自动等待channel的动作一直到channel被关闭
1. 无缓冲管道，有值写入管道则forrange读取值，否则阻塞直到管道关闭
2. 有缓存管道，缓冲中有值则forrange读取，否则阻塞直达管道关闭
 */
func main() {
    ch := make(chan int, 20)
    go func() {
        for i := 0; i < 10; i++ {
            ch <- i
        }

        close(ch)
    }()
    // 对于有缓冲管道，即使已经close掉了，forrange仍可以完整读取到写入的数据
    time.Sleep(10 * time.Second)

    for it := range ch {
       log.Println(">>>", it)
    }

    //close(ch)
    //time.Sleep(5 * time.Second)
}
