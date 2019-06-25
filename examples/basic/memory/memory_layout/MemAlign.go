package main

import (
	"fmt"
	"reflect"
)

/*
内存对齐

 */
type MyData struct {
	aByte 	byte
	//bByte 	byte
	aShort 	int16
	anInt 	int32
	aSlice 	[]byte
}

func main() {
	tp := reflect.TypeOf(MyData{})
	fmt.Printf("Struct is %d bytes long\n", tp.Size())

	n := tp.NumField()
	for i := 0; i < n; i++ {
		field := tp.Field(i)
		fmt.Printf("%s at offset %v, size=%d, align=%d\n", field.Name, field.Offset, field.Type.Size(),
			field.Type.Align())
	}
}
/*
output:
Struct is 32 bytes long
aByte at offset 0, size=1, align=1
aShort at offset 2, size=2, align=2
anInt at offset 4, size=4, align=4
aSlice at offset 8, size=24, align=8
 */