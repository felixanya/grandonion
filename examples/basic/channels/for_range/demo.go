package main

import (
    "log"
    "time"
)

/*
for range channel VS for select
1. for range channel 自动等待channel的动作一直到channel被关闭
2. for select
 */

func forRangeChan() {

}

func forSelectChan() {

}

func main() {
    ch := make(chan int, 20)
    go func() {
        for i := 0; i < 10; i++ {
            ch <- i
        }

        //close(ch)
        //time.Sleep(20 * time.Second)
    }()
    time.Sleep(1 * time.Second)
    close(ch)

    //for it := range ch {
    //    log.Println(">>>", it)
    //}

    timer := time.NewTimer(2 * time.Second)

OutFor:
    for {
        select {
        case it := <-ch:
            log.Println(">>>", it)
            timer.Reset(2 * time.Second)
        case <- timer.C:
            break OutFor
        }
    }

    //close(ch)
    time.Sleep(5 * time.Second)
}

