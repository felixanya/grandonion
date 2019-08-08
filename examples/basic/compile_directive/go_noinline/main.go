package main

//go:noinline
func appendStr(word string) string {
    return "new " + word
}

func main() {
    appendStr("world")
}
