package singleton

import (
	"fmt"
	"testing"
)

func TestGetSingleton(t *testing.T) {
	s := GetSingleton()
	fmt.Printf("s1: %p\n", s)

	s2 := GetSingleton()
	fmt.Printf("s2: %p\n", s2)
}
