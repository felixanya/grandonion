package main

import (
	"log"
	"time"
)

/*
令牌桶算法实现限流
 */

// 令牌桶
type RateLimit chan struct{}
// 获得令牌桶
func NewRateLimiter(limit int, rate time.Duration) RateLimit {
	r := make(RateLimit, limit)
	ticker := time.NewTicker(rate)
	// 按照传入速率往桶里放入令牌
	go func() {
		for {
			for i := 0; i < limit; i++ {
				select {
				case r <- struct{}{}:
				default:
				}
			}
			<-ticker.C
		}
	}()
	return r
}
// 获取令牌
func (r RateLimit) Limit(n int) {
	for i := 0; i < n; i++ {
		<-r
	}
}
func main() {
	start := time.Now()
	// 获得限速器 限制条件为 100/s
	rateLimit := NewRateLimiter(10, time.Second)
	// 模拟500个写入
	for i := 0; i < 500; i++ {
		rateLimit.Limit(1)
		log.Println("this is", i)
	}
	log.Println("use time", time.Now().Sub(start))
}
