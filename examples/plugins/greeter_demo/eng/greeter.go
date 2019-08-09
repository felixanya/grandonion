package eng

import "fmt"

type greeting string

func (g greeting) Greet() {
    fmt.Println("Hello Universe")
}

// exported as symbol named "Greeter"
var Greeter greeting

// 编译：go build -buildmode=plugin -o eng/eng.so eng/greeter.go