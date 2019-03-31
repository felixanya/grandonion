package main

import (
	"fmt"
	"github.com/aaronize/grandonion/examples/cycle_import/apkg"
	"github.com/aaronize/grandonion/examples/cycle_import/bpkg"
	"github.com/aaronize/grandonion/examples/cycle_import/common"
	"github.com/aaronize/grandonion/examples/cycle_import/cpkg"
	"github.com/aaronize/grandonion/examples/cycle_import/dpkg"
)

//func main1() {
//	// 报错：import cycle not allowed
//	apkg.Foo("apkg, good morning!")
//}

func main() {
	a := apkg.NewA()
	b := bpkg.NewB()
	a.Setb(b)
	fmt.Println(a.Foo("apkg interface implement, good morning!"))
	fmt.Println("----------------------------------------------")

	// 隐藏实现
	cc := cpkg.NewC()
	fmt.Println(cc.Deadline())
	fmt.Println(cc.Value(16))
	fmt.Println("----------------------------------------------")

	// 类似多态
	var dd common.D
	dd = dpkg.NewD()
	dd.SetKV("nihao", "shijie")
	fmt.Println(dd.GetValue("nihao"))
	fmt.Println("----------------------------------------------")

	bt := bpkg.NewBt("say", "hello")
	fmt.Println(bt.GetKey(), bt.GetVal())
	bt.SetKey("said")
	bt.SetVal("nihao")
	fmt.Println(bt.GetKey(), bt.GetVal())
}
