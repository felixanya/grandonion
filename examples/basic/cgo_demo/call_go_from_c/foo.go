package go_callback

/*
#include <stdio.h>
extern void ACFunction();
 */
import "C"
import "fmt"

//export AGoFunction
func AGoFunction() {
    fmt.Println("This is AGoFunction()")
}

func Example1() {
    C.ACFunction()
}
