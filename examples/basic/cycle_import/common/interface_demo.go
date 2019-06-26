package common

type Vari interface {

	ExecTask() chan<- map[string][]string

	Done() <-chan bool

	Err() error
}
