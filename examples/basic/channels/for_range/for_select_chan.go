package main

import (
    "log"
    "time"
)
/*

 */
func main() {
    ch := make(chan int, 20)
    go func() {
        for i := 0; i < 10; i++ {
            ch <- i
        }
        close(ch)
    }()

    timer := time.NewTimer(1 * time.Second)
OutFor:
    for {
        select {
        case it := <-ch:
            time.Sleep(1 * time.Second)
            // 读取完写入数据后，开始读取零值
            log.Println(">>>", it)
            timer.Reset(1 * time.Second)
        case <- timer.C:
            break OutFor
        }
    }
}
