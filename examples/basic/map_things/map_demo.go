package main

import "fmt"

func main() {
	mp := map[string]string{"1": "hello", "2": "world"}
	val, ok := mp["3"]
	fmt.Println("val: ", val)
	fmt.Println("ok: ", ok)
}
