package main

import (
	"fmt"
)

func main() {
	//runtime.GOMAXPROCS(1)
	GotoLabel()
	fmt.Println("-----------------------------")
	BreakToLabel()
	fmt.Println("-----------------------------")
	ContinueToLabel()
}

func GotoLabel() {
// 从label1开始重新执行循环
//label1:
	for i := 0; i < 10; i++ {
		if i == 3 {
			goto label2
		}
		//println("cycle:", i)
		fmt.Println("cycle:", i)
	}

// 从label2向下执行
label2:
	fmt.Println("end")
}

func BreakToLabel() {
label1:
	for i := 0; i < 10; i++ {
		if i == 3 {
			// 只能向上跳，直接跳出循环，不在执行label1下的循环
			break label1
		}
		//println("cycle:", i)
		fmt.Println("cycle:", i)
	}

// break跳不到这里
//label2:
	fmt.Println("end")
}

func ContinueToLabel() {
label1:
	for i := 0; i < 10; i++ {
		if i == 3 {
			// 只能向上跳，跳过此次循环
			continue label1
		}
		//println("cycle:", i)
		fmt.Println("cycle:", i)
	}

//label2:
	fmt.Println("end")
}
