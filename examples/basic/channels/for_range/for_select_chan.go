package main

import (
    "fmt"
    "log"
    "time"
)
/*
select关键字可用于多个channel的场景，这些channel会通过类似于are-you-ready polling的机制来工作。
1. 用于接收数据
2. 用于发送数据
 */
func main() {
    ch := make(chan int, 20)
    ch2 := make(chan string, 20)

    in := make(chan string)

    go func() {
        for i := 0; i < 10; i++ {
            time.Sleep(500 * time.Millisecond)
            ch <- i
        }
        //close(ch)
    }()
    go func() {
        for i := 0; i < 12; i++ {
            time.Sleep(50 * time.Millisecond)
            ch2 <- fmt.Sprintf("第%d个值", i)
            time.Sleep(300 * time.Millisecond)
        }
    }()
    go func() {
        for i := 0; i < 5; i++ {
            log.Println(">>>current timestamp ", <-in)
            time.Sleep(650 * time.Millisecond)
        }
    }()

    timer := time.NewTimer(1 * time.Second)
OutFor:
    for {
        select {
        case it := <-ch:
            // 读取完写入数据后，开始读取零值
            log.Println(">>>", it)
            timer.Reset(1 * time.Second)
        case st := <-ch2:
            log.Println(">>>", st)
            timer.Reset(1 * time.Second)
        case in <- time.Now().String():
            timer.Reset(1 * time.Second)
        case <- timer.C:
            break OutFor
        }
    }

    close(ch)
    close(ch2)
}
