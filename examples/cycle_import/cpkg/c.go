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
	// 这里要返回指针，否则会报错：
	// cannot use c (type CImpl) as type common.C in return argument
	// method has pointer receiver
	return &c
}

func (c *CImpl) Deadline() string {
	return "--- CImpl.Deadline---"
}

func (c *CImpl) Value(cv int) int {
	return cv * c.factor
}
