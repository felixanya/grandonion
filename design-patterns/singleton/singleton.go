package singleton

import "sync"

type Singleton struct {
}

var (
	singleton *Singleton
	once sync.Once
)

func GetSingleton() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})
	return singleton
}

// 懒汉式
//func GetSingleton2() *Singleton {
//	if singleton == nil {
//		return &Singleton{}
//	}
//	return singleton
//}


// 饿汉式
//var singleton3 = &Singleton{}
//
//func GetSingleton3() *Singleton {
//	return singleton3
//}

// 加锁
//var mu sync.Mutex
//
//func GetSingleton4() *Singleton {
//	if singleton == nil {
//		mu.Lock()
//		defer mu.Unlock()
//
//		if singleton == nil {
//			return &Singleton{}
//		}
//	}
//	return singleton
//}
