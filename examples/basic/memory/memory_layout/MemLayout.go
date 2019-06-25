package main

import (
	"fmt"
	"unsafe"
)

type MyData2 struct {
	aByte 	byte
	//bByte 	byte
	anInt 	int64
	aShort 	int16

	aSlice 	[]byte
}

func main() {
	data := MyData2{
		aByte: 0x1,
		aShort: 0x0203,
		anInt: 0x04050607,
		aSlice: []byte{0x08, 0x09, 0x0a},
	}

	dataBytes := (*[32]byte)(unsafe.Pointer(&data))
	//fmt.Printf("Bytes are %#v\n", dataBytes)
	fmt.Printf("Bytes are %#v\n", dataBytes)
}
