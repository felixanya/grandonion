package main

import (
    "fmt"
)

type Stringify interface {
    ToString() string
}

type String struct {
    data string
}

func (s *String) ToString() string {
    return s.data
}

func GetString() *String {
    return nil
}

func CheckString(s Stringify) bool {
    println(">>>>string in checkstring>>>>", s)
    return s == nil
}

func CheckNilInterface(i interface{}) bool {
    println(">>>check nil interface>>>", i)
    return i == nil
}

type T int

func main() {
    var sf Stringify
    s := GetString()
    println(">>>>>getstring>>>>", s)
    sf = s
    println(">>>>>interface type>>>>", sf)

    fmt.Println(CheckString(s))

    var str *String
    str = nil
    println("str==nil?: ", str == nil) // true

    var mock interface{}
    mock = str
    println("mock==nil?: ", mock == nil) // false
    println("mock==nil?: ", mock == (*String)(nil)) // true

    var t *String
    var i interface{} = t
    println("t: ", t)
    println("i: ", i)

    fmt.Println("CheckNilInterface: ", CheckNilInterface(nil)) // true
    fmt.Println("CheckNilInterface: ", CheckNilInterface(t)) // false
}
