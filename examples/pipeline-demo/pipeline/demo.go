package main

import "fmt"

/*
	producer()负责生产数据，它会把数据写入通道，并把它写数据的通道返回。
	square()负责从某个通道读数字，然后计算平方，将结果写入通道，并把它的输出通道返回。
	main()负责启动producer和square，并且还是消费者，读取suqre的结果，并打印出来。
 */
func producer(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()

	return out
}

func square(inChan <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range inChan {
			out <- n * n
		}
	}()

	return out
}

func main() {
	in := producer(1, 2, 3, 4, 5)
	ch := square(in)

	for ret := range ch {
		fmt.Printf("%3d", ret)
	}
	fmt.Println()
}
