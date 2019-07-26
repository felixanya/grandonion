package directory_operate

import (
    "fmt"
    "io/ioutil"
    "log"
)

func NewIoutilDir() string {
    // 缓存目录中创建一个test开头的目录（win10: C:\Users\ADMINI~1\AppData\Local\Temp\）
    path, err := ioutil.TempDir("", "test")
    if err != nil {
        log.Println(err)
        return ""
    }
    fmt.Println(path)
    return path
}
