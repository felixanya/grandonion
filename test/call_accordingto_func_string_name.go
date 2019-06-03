package main

import (
	"errors"
	"fmt"
	"time"
)

var funcNameMap map[string]interface{}

func fn1(s string) error {
	fmt.Println(">>>fn1>>>", s)
	return nil
}

func fn2(t int64, rc chan<- map[int] error) {
	time.Sleep(time.Duration(t) * time.Second)
	fmt.Println(">>>>>", t)
	if t % 2 == 0 {
		//return
		rc <- map[int]error{int(t): errors.New("非奇数错误")}
		return
	}
	rc <- map[int]error{int(t): nil}
	return
}

func main() {
	arrayList := []int64{6, 15, 12, 3, 9, 5}
	var retChan = make(chan map[int] error)

	for _, i := range arrayList{
		go fn2(int64(i), retChan)
	}

	for range arrayList {
		select {
		case mp := <-retChan:
			fmt.Println("输出结果：", mp)
		case <- time.After(4 * time.Second):
			fmt.Println("超时")
		}
	}

	fmt.Println("结束。")
}
