package main

import "fmt"

type T struct {
    Name    string
}

func (t T) Method1() {
}

func (t *T) Method2() {
    fmt.Println("Method2, ", t.Name)
}

func (t *T) Method3() {
}

func main() {
    //var t = T{}
    t := T{}
    t.Method1()
    // TODO 非*T类型应该不能调用receiver为*T类型的方法的啊？？？？
    t.Method2()
    t.Method3()

    var pt = &T{}
    pt.Method1()
    pt.Method2()
}
