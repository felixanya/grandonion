package main

import (
	"log"
	"runtime"
)

var integerMap map[int]int
var cnt = 8192

func main() {
	log.Println("进入主程序：")
	printMemStats()

	initMap()
	runtime.GC()
	log.Println("初始化map，并进行一次GC：")
	printMemStats()

	log.Println("执行delete函数删除前map长度：", len(integerMap))
	for i := 0; i < cnt; i++ {
		delete(integerMap, i)
	}
	log.Println("执行delete函数删除后的map长度:", len(integerMap))

	runtime.GC()
	log.Println("执行delete函数删除map元素并GC 后：")
	printMemStats()

	integerMap = nil
	runtime.GC()
	log.Println("map指向nil，并GC后：")
	printMemStats()
}

func initMap() {
	integerMap = make(map[int]int, cnt)

	for i := 0; i < cnt; i++ {
		integerMap[i] = i
	}
}

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("Alloc = %v TotalAlloc = %v Sys = %v NumGC = %v\n", m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024, m.NumGC)
}
