package main

import (
    "fmt"
    "github.com/aaronize/grandonion/examples/basic/compile_directive/go_linkname/apkg"
)

// `//go:linkname localname importpath.name`
//The `//go:linkname` directive instructs the compiler to use “importpath.name” as the object file symbol name
//for the variable or function declared as “localname” in the source code. Because this directive can subvert
//the type system and package modularity, it is only enabled in files that have imported "unsafe".

// 这种方式可以访问一个包中unexported的属性
// 这个编译指令破坏了类型系统和包的模块化，所有只能在unsafe的情况下才能使用

func main() {
    s := apkg.Say("world!")
    fmt.Println(s)
}
