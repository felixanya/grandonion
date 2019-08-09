package main

import "context"

type TimerTask struct {

}

func (t *TimerTask) Available() bool {

	return true
}

func (t *TimerTask) Enable() error {

	return nil
}

func (t *TimerTask) Disable() error {

	return nil
}

func (t *TimerTask) Execute(ctx context.Context, in ...interface{}) (interface{}, error) {

	return nil, nil
}

// export
var TaskPlugin *TimerTask
