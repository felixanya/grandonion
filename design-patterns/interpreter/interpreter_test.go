package interpreter

import (
	"fmt"
	"testing"
)

func TestNewTerminalExpression(t *testing.T) {
	expr1 := NewTerminalExpression("john")
	expr2 := NewTerminalExpression("robert")

	orDemo := NewOrExpression(expr1, expr2)
	orRet := orDemo.Interpret("john")
	fmt.Println("or expression result: ", orRet)
}
