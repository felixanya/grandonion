package main

import (
	"fmt"
	apkg2 "github.com/aaronize/grandonion/examples/basic/cycle_import/apkg"
	bpkg2 "github.com/aaronize/grandonion/examples/basic/cycle_import/bpkg"
	common2 "github.com/aaronize/grandonion/examples/basic/cycle_import/common"
	cpkg2 "github.com/aaronize/grandonion/examples/basic/cycle_import/cpkg"
	dpkg2 "github.com/aaronize/grandonion/examples/basic/cycle_import/dpkg"
)

//func main1() {
//	// 报错：import cycle not allowed
//	apkg.Foo("apkg, good morning!")
//}

func main() {
	a := apkg2.NewA()
	b := bpkg2.NewB()
	a.Setb(b)
	fmt.Println(a.Foo("apkg interface implement, good morning!"))
	fmt.Println("----------------------------------------------")

	// 隐藏实现
	cc := cpkg2.NewC()
	fmt.Println(cc.Deadline())
	fmt.Println(cc.Value(16))
	fmt.Println("----------------------------------------------")

	// 类似多态
	var dd common2.D
	dd = dpkg2.NewD()
	dd.SetKV("nihao", "shijie")
	fmt.Println(dd.GetValue("nihao"))
	fmt.Println("----------------------------------------------")

	bt := bpkg2.NewBt("say", "hello")
	fmt.Println(bt.GetKey(), bt.GetVal())
	bt.SetKey("said")
	bt.SetVal("nihao")
	fmt.Println(bt.GetKey(), bt.GetVal())
}
