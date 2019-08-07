package apkg

import _ "github.com/aaronize/grandonion/examples/basic/compile_directive/go_linkname/bpkg"

// 1. 需要引用下映射的包路径，告诉编译器去哪个包下找映射的函数
// 2. 在当前包下创建任意一个汇编文件(.s)，来绕过编译器的空函数体检查
func Say(name string) string

func Greet(name string) string {
	return Say(name)
}
