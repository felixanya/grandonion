package main

import (
	"fmt"
	"strings"
)

func main() {
	order := "-status"
	fmt.Println(strings.HasPrefix(order, "-"))
	order = strings.TrimPrefix(order, "-")

	fmt.Println(order)
}
