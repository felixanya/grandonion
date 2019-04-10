package Observer

import "fmt"

// 参考：https://blog.csdn.net/jeanphorn/article/details/78784197

type IObserver interface {
	Notify()
}

type Observer struct {
}

func (o *Observer) Notify() {
	fmt.Println("触发观察者...")
}
