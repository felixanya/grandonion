package main

import (
    "fmt"
    "time"
)

const (
    _ = iota
    RoomPlan
    ResDeploy
    NetConf
    ServerDeploy
    ServerClean
    RoomExtends

    RoomPlanOne = 101
    RoomPlanTwo = iota

    ServerCleanOne
)

func main() {
    fmt.Println(RoomPlan)
    fmt.Println(RoomExtends)

    fmt.Println(RoomPlanOne)
    fmt.Println(RoomPlanTwo)

    fmt.Println(ServerCleanOne)

    time.Sleep(2 * time.Second)
}
