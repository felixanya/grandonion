package main

import (
    "fmt"
    "log"
    "path/filepath"
)

func main() {
    path := "/var/lib/project/node.db/"
    fmt.Println(">>>", filepath.Dir(path)) // 提取目录
    fmt.Println(">>>", filepath.Base(path)) // 最底层目录或文件

    pathAbs, err := filepath.Abs(path) // 绝对路径
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(">>>", pathAbs)
    fmt.Println(">>>", filepath.Ext(path)) // 文件后缀
    fmt.Println(">>>", filepath.Join(path, "/node.log")) // 拼接
    fileNames, err := filepath.Glob(path + "*") // 返回和pattern匹配的所有文件名
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(">>>", fileNames)
}
