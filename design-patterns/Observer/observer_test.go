package Observer

import (
	"fmt"
	"sync"
	"testing"
)

func TestObserver_Notify(t *testing.T) {
	sub := &Subject{
		observers: make(map[IObserver]struct{}),
	}
	ob1 := &Observer{Id: 1}
	ob2 := &Observer{Id: 2}
	ob3 := &Observer{Id: 3}
	ob4 := &Observer{Id: 4}

	sub.Register(ob1, ob2, ob3, ob4)

	e := &Event{Data: "hello", wg: sync.WaitGroup{}}

	sub.Notify(e)
	fmt.Println("----------------")

	sub.Dismiss(ob3)
	e2 := &Event{Data: "nihao", wg: sync.WaitGroup{}}
	sub.Notify(e2)
}
