package decorator_generic

import (
	"fmt"
	"testing"
)

func foo(a, b, c int) int {
	fmt.Printf("%d, %d, %d \n", a, b, c)
	return a + b + c
}

func bar(a, b string) string {
	fmt.Printf("%s, %s \n", a, b)
	return a + b
}

type MyFoo func(int, int, int) int
var myfoo MyFoo

func TestDecorator(t *testing.T) {
	Decorator(&myfoo, foo)
	myfoo(1, 2, 3)
}
