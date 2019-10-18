package main

import "fmt"

func main() {
    mp := map[string]interface{}{"UUID": "abcdefghijklmn", "Name": "aaron.test"}

    if mp["nihao"] == "" || mp["nihao"] == nil {
        fmt.Println("UUID 为空")
    }
}
