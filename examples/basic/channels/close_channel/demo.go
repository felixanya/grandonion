package main

import "fmt"

/*
After the last value has been received from a closed channel c,
any receive from c will succeed without blocking,
returning the zero value for the channel element.
The form
	x, ok := <-c
will also set ok to false for a closed channel.
 */

func main() {
    ch := make(chan int, 5)

    ch <- 1
    ch <- 2

    close(ch)
    fmt.Println(">>>", <-ch) // >>> 1
    rv, ok := <-ch
    fmt.Printf(">>>read val: %d, ok: %t\n", rv, ok) // >>>read val: 2, ok: true
    // channel关闭后，仍然可以读取已写入的数据。读取完后开始读取零值。
    fmt.Println(">>>", <-ch) // >>> 0
    rv, ok = <-ch
    fmt.Printf(">>>read val: %d, ok: %t\n", rv, ok) // >>>read val: 0, ok: false

    // channel关闭后继续写入则会报错
    // ch <- 3 // panic: send on closed channel
}
