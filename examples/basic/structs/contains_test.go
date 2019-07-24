package structs

import (
    "fmt"
    "testing"
)

func TestConcreteFlow(t *testing.T) {
    cf := ConcreteFlow{}
    _ = cf.Add()
}

func TestNewConcreteFlow(t *testing.T) {
    cf := NewConcreteFlow(0, 0, "test_flow_name")
    _ = cf.Add()
    _ = cf.Update()
    fmt.Println(">>>", cf.FlowName)

    //ch := make(chan map[string]string, 5)
    //ret := make(map[string]string)
    //ret["key1"] = "val1"
    //ch<-ret
    //
    //ret2 := make(map[string]string)
    //ret2["key2"] = "val2"
    //ch <- ret2
    //
    //ret3 := make(map[string]string)
    //ret3["key3"] = "val3"
    //ch <- ret3
    //
    //retList := []map[string]string{ret, ret2, ret3}
    //
    //for range retList {
    //    for k, v := range <- ch {
    //        fmt.Printf(">>> %s:%s \n", k, v)
    //    }
    //}
}
