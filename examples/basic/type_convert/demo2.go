package main

import "fmt"

type CusError struct {

}

func (ce *CusError) Error() string {
    return "custom error"
}

func main() {
    err := Process()
    fmt.Println(err) // custom error

    println(err) // (0x4de080,0x0) 动态类型是：*CusError，动态值是0

    fmt.Println(err == nil) // false
    fmt.Println(err == (*CusError)(nil)) // true

    err = Process2()
    fmt.Println(err) // custom error

    println(err) // 0x0

    fmt.Println(err == nil) // true
    fmt.Println(err == (*CusError)(nil)) // true
}

func Process() error {
    var err *CusError = nil
    return err
}

//
func Process2() *CusError {
    var err *CusError = nil
    return err
}
