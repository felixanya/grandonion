package cgo_demo

/*
#include <stdio.h>
#include <stdlib.h>

void myprint(char* s)
{
    printf("%s\n", s);
}
 */
import "C"
// cgo代码块必须紧挨着import "C"

import "unsafe"

func Example1() {
    cs := C.CString("hello world from stdio\n")
    C.myprint(cs)
    C.free(unsafe.Pointer(cs))
}
