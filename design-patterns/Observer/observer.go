package Observer

import "fmt"

type IObserver interface {
	Notify()
}

type Observer struct {
}

func (o *Observer) Notify() {
	fmt.Println("触发观察者...")
}
