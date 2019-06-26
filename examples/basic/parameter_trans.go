package main

import "fmt"
/*
参数传递
 */
type ParamTrans struct {
	Name 		string
	Type 		string
}

func main() {
	pt := ParamTrans{Name: "aaron", Type: "type0"}
	fmt.Printf("before trans: %p\n", &pt)
	another(pt)
	anotherWithPointer(&pt)

	fmt.Printf("changes: %+v\n", pt) // name 没改变；type 被修改了
}

// 非指针参数传递，发生拷贝
func another(pt ParamTrans) {
	fmt.Printf("no pointer: %p\n", &pt)
	pt.Name = "changed name"
	return
}

func anotherWithPointer(pt *ParamTrans) {
	fmt.Printf("with pointer: %p\n", pt)
	pt.Type = "changed type"
	return
}
