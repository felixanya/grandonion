package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "runtime"
)

func main() {
    // 获取程序运行绝对路径
    dir, err := filepath.Abs(os.Args[0])
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("--->", dir)
    // 获取运行程序所在目录
    root := filepath.Dir(dir)

    funcName, file, line, ok := runtime.Caller(0)
    if ok {
        fmt.Println(runtime.FuncForPC(funcName).Name())
        file, _ = filepath.Rel(root, file)
        fmt.Println(file)
        fmt.Println(line)
    }
}
