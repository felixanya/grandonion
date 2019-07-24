package main

import (
    "fmt"
    "reflect"
)

type Flow interface {
    Add() error
}

type OfficalFlow struct {

}

func (of *OfficalFlow) Add() error {
    fmt.Printf(">>> Add In OfficalFlow>>>")
    return nil
}

func Checking(beans ...interface{}) error {
    for _, bean := range beans {
        switch bean.(type) {
        case map[string]interface{}:
        default:
            sliceValue := reflect.Indirect(reflect.ValueOf(bean))
            sliceValue.MethodByName("Add").Call(nil)
        }
    }
    return nil
}

func main() {
    of := OfficalFlow{}
    _ = Checking(&of)
}
