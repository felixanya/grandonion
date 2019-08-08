package main

func main() {
    a := fn1()
    *a++
}

//go:noinline
func fn1() *int {
    i := 1
    return &i
}

// $ go build -gcflags "-m" mem_escape.go
//
// # command-line-arguments
// ./mem_escape.go:8:6: can inline fn1
// ./mem_escape.go:3:6: can inline main
// ./mem_escape.go:4:13: inlining call to fn1
// ./mem_escape.go:4:13: main &i does not escape
// ./mem_escape.go:10:12: &i escapes to heap
// ./mem_escape.go:9:5: moved to heap: i

// 加了noinline指令后
// $ go build -gcflags "-m" mem_escape.go
// # command-line-arguments
// ./mem_escape.go:11:12: &i escapes to heap
// ./mem_escape.go:10:5: moved to heap: i