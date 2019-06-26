package apkg

import (
	common2 "github.com/aaronize/grandonion/examples/basic/cycle_import/common"
	"strings"
)

type AImpl struct {
	b common2.B
}

func NewA() *AImpl {
	return new(AImpl)
}

func (a *AImpl) Setb(b common2.B) {
	a.b = b
}

func (a *AImpl) Foo(as string) string {
	return a.b.Add(as)
}

func (a *AImpl) Minus(as string) string {
	return strings.Trim(as, "\t")
}
