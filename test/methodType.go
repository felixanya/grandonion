package main

import "fmt"

type MType struct {
	a int
}

func (mt MType) Get() int {
	return mt.a
}

func (mt *MType) Set(a int) int {
	mt.a = a
	return mt.a
}

func main() {
	var mt MType
	fmt.Println(MType.Get(mt))
	fmt.Println((*MType).Set(&mt, 2))
}
