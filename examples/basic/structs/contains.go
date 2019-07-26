package structs

import (
    "fmt"
    "time"
)

type Flow interface {
    Add() error
    Del([]string) error
    Update() error
    Query() ([]Flow, error)
}

type BaseFlow struct {
    Id              int
    FlowType        int
    FlowName        string
    Status          int
    CreateAt        string
    UpdateAt        string
    LastUpdateUser  string
    DeleteAt        string
    Remark          string
}

func (bf *BaseFlow) Add() error {
    fmt.Println("Add In BaseFlow")
    return nil
}

func (bf *BaseFlow) Del(idList []string) error {
    fmt.Println("Del In BaseFlow")
    return nil
}

func (bf *BaseFlow) Update() error {
    fmt.Println("Update In BaseFlow")
    return nil
}

func (bf *BaseFlow) Query() ([]Flow, error) {
    fmt.Println("Query In BaseFlow")
    return nil, nil
}

//
type ConcreteFlow struct {
    BaseFlow

    DataList    []string
    TaskId      string
}

func NewConcreteFlow(flowType, status int, flowName string) *ConcreteFlow {
    return &ConcreteFlow{
        BaseFlow{
            FlowType: flowType,
            Status: status,
            FlowName: flowName,
            CreateAt: time.Now().Format("2006-01-02 15:04:05.999999"),
            UpdateAt: time.Now().Format("2006-01-02 15:04:05.999999"),
        },
        []string{"127.0.0.1", "192.168.168.123"},
        "123456",
    }
}

func (cf *ConcreteFlow) Add() error {
    fmt.Printf("Add In ConcreteFlow. Val: %+v\n", cf)
    return nil
}
