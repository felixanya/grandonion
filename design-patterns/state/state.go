package state

import "fmt"

type IState interface {
	doAction(*Context)
	toString() string
}

// context
type Context struct {
	state 	IState
}

func NewContext() *Context {
	return &Context{
		state: nil,
	}
}

func (c *Context) setState(state IState) {
	c.state = state
}

func (c *Context) getState() IState {
	return c.state
}

// start state
type StartState struct {

}

func (ss *StartState) doAction(context *Context) {
	fmt.Println("Player is in start state.")
	context.setState(ss)
}

func (ss *StartState) toString() string {
	return "start state"
}

// stop state
type StopState struct {

}

func (ss *StopState) doAction(context *Context) {
	fmt.Println("Player is in stop state.")
	context.setState(ss)
}

func (ss *StopState) toString() string {
	return "stop state"
}
