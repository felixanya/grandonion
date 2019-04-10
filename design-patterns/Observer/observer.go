package Observer

import "fmt"

// 参考：https://blog.csdn.net/jeanphorn/article/details/78784197

type IObserver interface {
	Update(*Event)
}

type Observer struct {
	Id 	int
}

func (o *Observer) Update(event *Event) {
	fmt.Printf("触发观察者%d号执行%s \n", o.Id, event.Data)
	event.wg.Done()
}
