package main

import (
	"fmt"
	"strings"
)

type channel struct {
	read  chan interface{}
	write chan interface{}
	done  chan bool
}

func main() {
	c := channel{make(chan interface{}), make(chan interface{}), make(chan bool)}
	go func() {
		var value interface{}
		for {
			select {
			case c.read <- value:
			case value = <-c.write:
				fmt.Println(">>", value)
			case <-c.done:
				fmt.Println("break")
				// break只会退出select的当前分支，并不能退出for循环
				break
			}
		}
	}()
	p := c.write
	p <- strings.Title("the nice world") //Title
	p <- strings.Contains((<-c.read).(string), "Nice")
	p <- strings.ContainsAny("The Nice World", "Nice")
	p <- strings.Count("string", "s")
	p <- strings.EqualFold("the", "The")
	p <- strings.HasPrefix("popar fish", "po")

	c.done <- true
	p <- strings.HasSuffix("poplar fish", "fish")
	p <- strings.HasSuffix("poplar fish", "fish")
	p <- strings.HasSuffix("poplar fish", "fish")
	p <- strings.HasSuffix("poplar fish", "fish")

}