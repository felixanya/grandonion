package http_net

import (
    "fmt"
    "log"
    "testing"
)

func TestAddBusinessTypePatch(t *testing.T) {
    s := &ServerInfo{
        Data: []map[string]interface{}{
            {"Usage": "uops", "SecUsage": "executor"},
        },
    }

    if err := AddBusinessTypePatch(s); err != nil {
        log.Println(err)
    }

    fmt.Printf("---> %+v", s)
}
