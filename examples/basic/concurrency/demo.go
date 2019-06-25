package main

import (
	"fmt"
	"runtime"
)

func main(){
	runtime.GOMAXPROCS(1)
	go func(){
		fmt.Println("hello world")
	}()
	go func(){
		//for {
		//}
	}()

	select {}
}

