package bpkg

import (
	common2 "github.com/aaronize/grandonion/examples/basic/cycle_import/common"
)

type BImpl struct {
	a common2.A
}

func NewB() *BImpl {
	return new(BImpl)
}

func (b *BImpl) Seta(a common2.A) {
	b.a = a
}

func (b *BImpl) Goo(bs string) string {
	return b.a.Minus(bs)
}

func (b *BImpl) Add(bs string) string {
	return bs + "---BImpl.Add---"
}
