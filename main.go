package main

import "fmt"

func main() {
	// 直接运行go run main.go 会提示.\main.go:6:14: undefined: Version错误，
	fmt.Println(Version)
}
