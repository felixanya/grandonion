package main

import (
	"fmt"
	"time"
)

func main() {
	var x, y int
	go func() {
		x = 1 // A1
		fmt.Print("y:", y, " ") // A2
	}()
	go func() {
		y = 1                   // B1
		fmt.Print("x:", x, " ") // B2
	}()

	time.Sleep(1 * time.Second)
}
/*
实际运行结果，会出现各种情况：
y:0 x:1
x:0 y:1
x:1 y:1
y:1 x:1
甚至
x:0 y:0
y:0 x:0
 */