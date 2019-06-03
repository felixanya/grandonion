package main

import (
	"fmt"
)

type student struct {
	Name string
	Age  int
}

func pase_student() map[string]*student {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		fmt.Printf("%p\n", &stu)
		m[stu.Name] = &stu
	}
	return m
}

/*
在`for range`迭代map或slice的语句中有两个变量`k`和`v`，其中`v`变量用于保存迭代过程中元素的值，
但是`v`只被声明了一次，此后都是将迭代出的值赋值给`v`，也即`v`变量的内存地址始终未变。
 */
func main() {
	students := pase_student()

	for k, v := range students {
		fmt.Printf("key=%s,value=%v \n", k, v)
	}
}
