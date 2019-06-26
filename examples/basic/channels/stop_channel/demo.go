package main

import (
	"fmt"
	"time"
)

func worker(stopCh <-chan interface{}) {
	go func() {
		defer fmt.Println("worker exit")

		t := time.NewTicker(time.Millisecond * 5000)

		// Using stop channel explicit exit
	//BREAK_LABEL_DEMO:
		//GOTO_LABEL_DEMO:
		for {
			fmt.Println("in for ...")
			select {
			case a := <-stopCh:
				fmt.Println("Recv stop signal.", a)
				// 1. return 直接退出函数
				//return
				// 2. break 只能退出select分支，继续执行for的下一次循环; 这里在stopCh显式关闭后，会一直取出零值，因此会一直执行这个分支
				break
				// 3. break label 跳出多层嵌套循环, 只能写在需要跳出的多重循环的前面
				//break BREAK_LABEL_DEMO
				// 4. 使用goto label跳转到标志位处，前面和后面都可以跳转到
				//goto GOTO_LABEL_DEMO
			case <-t.C:
				fmt.Println("Working .")
			}
		}

//GOTO_LABEL_DEMO:
	fmt.Println("start")
	<- t.C
	fmt.Println("end")

	}()
	return
}

func main() {

	stopCh := make(chan interface{})
	worker(stopCh)

	time.Sleep(time.Second * 10)
	// 关闭管道后，再从该管道取值，会取不断出空值（也即，关闭后就不会阻塞了）
	close(stopCh)

	// Wait some print
	time.Sleep(30 * time.Second)
	fmt.Println("main exit")
}