package bpkg

import "github.com/aaronize/grandonion/examples/cycle_import/common"

type BImpl struct {
	a 	common.A
}

func NewB() *BImpl {
	return new(BImpl)
}

func (b *BImpl) Seta(a common.A) {
	b.a = a
}

func (b *BImpl) Goo(bs string) string {
	return b.a.Minus(bs)
}

func (b *BImpl) Add(bs string) string {
	return bs + "---BImpl.Add---"
}
