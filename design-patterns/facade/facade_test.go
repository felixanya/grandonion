package facade

import "testing"

func TestNewDataHandler(t *testing.T) {
	dataHandler := NewDataHandler()
	dataHandler.Decode("hellowrold")
	dataHandler.Aggregate("helloworld")
	dataHandler.Encode("helloworld")
}
