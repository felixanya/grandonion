package main

import (
	"bytes"
	"sync"
)

type SafeBuffer struct {
	b 	bytes.Buffer
	rw 	sync.RWMutex
}

func (sb *SafeBuffer) Read(p []byte) (n int, err error) {
	sb.rw.RLock()
	defer sb.rw.Unlock()

	return sb.b.Read(p)
}

func (sb *SafeBuffer) Write(p []byte) (n int, err error) {
	sb.rw.Lock()
	defer sb.rw.Unlock()

	return sb.b.Write(p)
}
