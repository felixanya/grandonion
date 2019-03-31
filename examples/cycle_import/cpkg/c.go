package cpkg

import "github.com/aaronize/grandonion/examples/cycle_import/common"

type CImpl struct {
	common.C

	factor 	int
}

func NewC() common.C {
	c := CImpl{
		//C: ci,
		factor: 2,
	}
	// 这里要返回指针（返回值还是指针要和CImpl方法的receiver保持一致），否则会报错：
	// cannot use c (type CImpl) as type common.C in return argument:
	// CImpl does not implement common.C (Deadline method has pointer receiver)
	// 这是因为Deadline方法是pointer method，如果返回的receiver是值类型的话就无法调用。
	// value method可以被pointer和value对象调用，而pointer method只能被pointer对象调用。
	return &c
}

/*
pointer method可以修改对象的值，而value method不会，
所以如果在一个value对象上调用pointer method，编译器会对原来的值做一份拷贝，
并在拷贝后的值上执行函数，如果函数有修改原receiver的值，则修改的行为都发生在拷贝的值上，而不会影响原值，
这个错误很隐蔽，非常不容易调试发现，因此GO决定放弃这个错误的发生的可能性，
直接不支持pointer method被value对象调用。
 */
func (c *CImpl) Deadline() string {
	return "--- CImpl.Deadline---"
}

func (c *CImpl) Value(cv int) int {
	return cv * c.factor
}
