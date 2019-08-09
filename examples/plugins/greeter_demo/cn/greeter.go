package cn

import "fmt"

type greeting string

func (g greeting) Greet() {
    fmt.Println("你好, 世界")
}

// exported as symbol named "Greeter"
var Greeter greeting

// 编译：go build -buildmode=plugin -o cn/cn.so cn/greeter.go