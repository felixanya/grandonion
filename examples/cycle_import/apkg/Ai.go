package apkg

import (
	"github.com/aaronize/grandonion/examples/cycle_import/common"
	"strings"
)

type AImpl struct {
	b 	common.B
}

func NewA() *AImpl {
	return new(AImpl)
}

func (a *AImpl) Setb(b common.B) {
	a.b = b
}

func (a *AImpl) Foo(as string) string {
	return a.b.Add(as)
}

func (a *AImpl) Minus(as string) string {
	return strings.Trim(as, "\t")
}
