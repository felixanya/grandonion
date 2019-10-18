package structs

import (
    "fmt"
    "testing"
)

func TestComputer_AddRemark(t *testing.T) {
    c := NewComputer()
    c2 := Computer{}

    _ = c.AddRemark("")
    _ = c2.AddRemark("")

    //fmt.Println(c.GetProductNo())   // runtime error: invalid memory address or nil pointer dereference
    fmt.Println(c2.GetProductNo())
}

/*
receiver类型为指针的方法
 */
