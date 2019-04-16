package decorator

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

// https://coolshell.cn/articles/17929.html

type SumFunc func(int64, int64) int64

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func timedSumFunc(fn SumFunc) SumFunc {
	return func(start int64, end int64) int64 {
		defer func(t time.Time) {
			fmt.Printf("--- Time Elapsed (%s): %v --- \n", getFunctionName(fn), time.Since(t))
		}(time.Now())
		return fn(start, end)
	}
}

func Sum1(start, end int64) int64 {
	var sum int64 = 0
	if start > end {
		start, end = end, start
	}
	for i := start; i <= end; i++ {
		sum += i
	}
	return sum
}

func Sum2(start, end int64) int64 {
	if start >= end {
		start, end = end, start
	}
	return (end - start + 1) * (end + start) / 2
}
