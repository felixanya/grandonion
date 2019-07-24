package directory_operate

import (
    "fmt"
    "testing"
    "time"
)

func TestNewIoutilDir(t *testing.T) {
    path := NewIoutilDir()
    fmt.Println(path)
    fmt.Println(">>>", fmt.Sprintf("清洗流程-%s-%s", "HD06", time.Now().Format("20060102150405")))
    fmt.Println(time.Now().Format("20060102150405"))
}
