package Observer

import "sync"

type Event struct {
	Data 	string
	wg 		sync.WaitGroup
}
