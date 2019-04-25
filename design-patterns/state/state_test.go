package state

import (
	"fmt"
	"testing"
)

func TestNewContext(t *testing.T) {
	context := NewContext()
	startState := new(StartState)
	startState.doAction(context)

	fmt.Println(context.getState().toString())

	stopState := new(StopState)
	stopState.doAction(context)

	fmt.Println(context.getState().toString())
}
