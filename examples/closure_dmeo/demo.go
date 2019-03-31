package main

import (
	"fmt"
	"time"
)

func main() {
	var m = [...]int{1, 2, 3, 4, 5}

	for i, v := range m {
		go func() {
			time.Sleep(1 * time.Second)
			// 打印的i, v的值都是4, 5（也有可能在最早的几个goroutine打印的是前面的值，看循环和goroutine运行的时间）
			fmt.Println("test1:", i, v)
		}()
		//time.Sleep(1 * time.Second)

		go func(a, b int) {
			// 把每次循环的i, v的值传到闭包中保存，可以避免上面的问题，但顺序也是不能保证的
			time.Sleep(1 * time.Second)
			fmt.Println("test1:", a, b)
		}(i, v)
	}

	time.Sleep(10 * time.Second)
}
