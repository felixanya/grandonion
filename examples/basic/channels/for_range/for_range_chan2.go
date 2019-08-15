package main

import (
    "log"
    "time"
)

func main() {
    ch := make(chan int, 20)

    if err := sendTask(ch); err != nil {
        log.Fatal(err)
    }

    time.Sleep(10 * time.Second)
    if err := collectResult(ch); err != nil {
        log.Fatal(err)
    }

    return
}

func sendTask(ch chan<- int) error {
    for i := 0; i < 19; i++ {
        ch <- i
    }
    // 这里关闭后collectResult中仍然可以读取到ch中的数据
    // 如果这里（或者在forrange之前的其他地方）不关闭ch的话，forrange将会等待，最终造成死锁
    close(ch)

    return nil
}

func collectResult(ch <-chan int) error {
    for i := range ch {
        log.Println(">>> ", i)
    }
    return nil
}
