package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a byte = 'a'
	fmt.Printf("%d %T\n", a, a)

	var b rune = '你'
	fmt.Printf("%d %T\n", b, b)

	var c int = 'f'
	fmt.Printf("%d %T\n", c, c)

	var timestamp_uint32 uint64 = 9559267382
	var timestamp_int int = 1559267382
	var timestamp_byte int = 15
	var timestamp_bye_array []byte = []byte{15}
	fmt.Printf("%s\n", string(timestamp_uint32)) // 乱码
	fmt.Printf("%s\n", string(timestamp_int)) // 乱码
	fmt.Printf("%s\n", string(timestamp_byte)) // 乱码
	fmt.Printf("%s\n", string(timestamp_bye_array)) // 乱码
	fmt.Printf("%s\n", string(b)) // 你

	fmt.Printf("%s\n", strconv.Itoa(int(timestamp_uint32)))
}
