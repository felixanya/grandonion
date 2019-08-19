package agent_plugin

import (
    "fmt"
    "log"
    "testing"
)

func TestNewScanner(t *testing.T) {
    scanner := NewScanner("./plugins")
    all, err := scanner.ListAllPlugin()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("list all plugin: %+v\n", all)

    available, err := scanner.ListAvailablePlugin()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("list all available plugin: %+v\n", available)


}
