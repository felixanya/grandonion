package main

import "fmt"

/*
go run/build时编译器只会加载命令后指定的文件，即使有依赖关系也不会主动加载main包下的其他文件。
因此，在编译的时候需要显式的列出所有需要编译的文件，否则会提示undefined

例，
    go build -o main version.go main.go 或 go build -o main *.go
 */
func main() {
    fmt.Printf("version: %s\ncommit: %s", version, commit)
}
