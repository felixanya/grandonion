package main

import "fmt"

func main() {
	mp := make(map[string]string)
	mp["ni"] = "hao"
	for k, v := range mp {
		fmt.Println(k, v)
	}
}
