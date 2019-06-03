package main

import (
	"fmt"
	"strings"
)

func main() {
	s := " 1000\n"
	s1 := strings.Trim(strings.Trim(s, "\n"), " ")

	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println(s1)
}
