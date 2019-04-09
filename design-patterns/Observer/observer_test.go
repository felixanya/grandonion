package Observer

import (
	"testing"
)

func TestObserver_Notify(t *testing.T) {
	sub := &Subject{}
	obs := &Observer{}

	sub.AddObservers(obs)

	sub.NotifyObservers()
}
