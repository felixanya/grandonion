package main

import (
    "fmt"
)

func main() {
    breakForSelect()
    fmt.Println("_________________")
    gotoForSelect()
    fmt.Println(">>>>>>>>>>>>>>>>>")
    continueForSelect()
}

func breakForSelect() {
    ch := make(chan int, 10)
    go func() {
        for i := 1; i < 11; i++ {
            ch <- i
        }
    }()

OuterFor:
    for {
        select {
        case it := <-ch:
            fmt.Println(">>>", it)
            if it == 5 {
                // 只有break只是跳出了select，并没有跳出for，---for tail---还是会打印
                //break
                // 跳出OuterFor标记的这一层循环。标签指定一层for循环，因此只能紧挨着for开头定义。
                break OuterFor
            }
            fmt.Println(">>>tail>>>")
        }
        fmt.Println("---for tail---")
    }
    close(ch)
}

func gotoForSelect() {
    ch := make(chan int, 10)
    go func() {
        for i := 1; i < 11; i++ {
            ch <- i
        }
    }()

//Forehead:
    for {
        select {
        case it := <-ch:
            fmt.Println(">>>", it)
            if it == 5 {
                // 跳到Forehead标签处，重新执行for
                //goto Forehead

                // 跳到Behind标签处，接着向后执行
                goto Behind
            }
            fmt.Println(">>>tail>>>")
        }
    }
Behind:

    close(ch)
}

func continueForSelect() {
    ch := make(chan int, 10)
    go func() {
        for i := 1; i < 11; i++ {
            ch <- i
        }
    }()

OuterFor:
    for {
        select {
        case it := <-ch:
            fmt.Println(">>>", it)
            if it == 5 {
                // continue 是for的continue，跳出此次循环，可以看到---for tail---也被跳过
                //continue
                // 用法和break类似
                continue OuterFor
            }
            fmt.Println(">>>tail>>>")
        }
        fmt.Println("---for tail---")
    }
    close(ch)
}
