package main

import (
	"fmt"
	"sync"
	"time"
)

type UserAges struct {
	ages map[string]int
	//sync.Mutex
	// promote1：互斥锁换成读写锁
	sync.RWMutex
}
func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()

	ua.ages[name] = age
}
func (ua *UserAges) Get(name string) int {
	// promote2：加锁读锁
	ua.RLock()
	defer ua.RUnlock()

	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func main() {
	// 设置使用单核时，可以避免错误fatal error: concurrent map read and map write
	//runtime.GOMAXPROCS(1)
	ua := &UserAges{ages: make(map[string]int)}
	//ua.Add("aaron", 20)
	//ua.Add("jeseling", 18)
	for i := 0; i < 1000; i++ {
		go func(age int) {
			ua.Add("aaron", age)
			//fmt.Println(ua.Get("aaron"))
		}(i)
		go func() {
			//ua.Add("aaron", i - 2)
			fmt.Println(ua.Get("aaron"))
		}()
	}

	time.Sleep(10 * time.Second)
}
