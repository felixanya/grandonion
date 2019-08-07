package internal

import "fmt"

// 用于测试外部导入internal包下属性
func MyPrint2() {
    fmt.Println("Can the function in internal pkg be imported?")
}
