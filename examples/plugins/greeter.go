package main

//import "plugin"; ...
import (
    "fmt"
    "log"
    "os"
    "plugin"
)

var defaultLang = "eng"

type Greeter interface {
    Greet()
}

func main() {
    lang := defaultLang

    if len(os.Args) == 2 {
        lang = os.Args[1]
    } else {
        log.Fatal("参数错误")
    }

    plg, err := plugin.Open(fmt.Sprintf("./lang/%s.so", lang))
    if err != nil {
        log.Fatal(err)
    }

    greeterSymbol, err := plg.Lookup("Greeter")
    if err != nil {
        log.Fatal(err)
    }

    var greeter Greeter
    greeter, ok := greeterSymbol.(Greeter)
    if !ok {
        log.Fatal("get greeter fail")
    }

    greeter.Greet()
}

//[root@centos plugins]# go run greeter.go eng
//Hello Universe
//[root@centos plugins]# go run greeter.go cn
//你好, 世界
