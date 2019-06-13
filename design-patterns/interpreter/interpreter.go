package interpreter

import "strings"

type Expression interface {
	Interpret(string) bool
}

// terminal
type TerminalExpression struct {
	data 	string
}

func NewTerminalExpression(data string) *TerminalExpression {
	return &TerminalExpression{
		data: data,
	}
}

func (ti *TerminalExpression) Interpret(context string) bool {
	if strings.Contains(context, ti.data) {
		return true
	}
	return false
}

// or
type OrExpression struct {
	expr1 		Expression
	expr2 		Expression
}

func NewOrExpression(expr1, expr2 Expression) *OrExpression {
	return &OrExpression{
		expr1: expr1,
		expr2: expr2,
	}
}

func (oe *OrExpression) Interpret(context string) bool {
	return oe.expr1.Interpret(context) || oe.expr2.Interpret(context)
}

// and
type AndExpression struct {
	expr1 		Expression
	expr2 		Expression
}

func NewAndExpression(expr1, expr2 Expression) *OrExpression {
	return &OrExpression{
		expr1: expr1,
		expr2: expr2,
	}
}

func (ae *AndExpression) Interpret(context string) bool {
	return ae.expr1.Interpret(context) && ae.expr2.Interpret(context)
}
