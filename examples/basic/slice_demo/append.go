package main

import "fmt"

func main() {
	// make会初始化cap=10个零值
	s := make([]int, 0, 2)
	t := []int{2, 1, 3, 4, 5, 6, 7}

	fmt.Println("1: ", s)
	fmt.Println("length:", len(s))
	fmt.Println("capacity:", cap(s))

	// capacity 扩容策略：
	// 第一次倍增，后面每次缓慢增加
	s = append(s, t...)
	fmt.Println("2：", s)
	fmt.Println("length:", len(s))
	fmt.Println("capacity:", cap(s))
}

/*
output:
1:  [0 0 0 0 0 0 0 0 0 0]
2： [0 0 0 0 0 0 0 0 0 0 21 22 23 24 25]
 */
