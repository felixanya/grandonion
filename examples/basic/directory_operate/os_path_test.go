package directory_operate

import (
    "fmt"
    "testing"
)

func TestDirOperations(t *testing.T) {
    dir, _ := DirOperations()
    fmt.Println(dir)
}
