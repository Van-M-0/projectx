package util

import (
	"sync"
)

type WaitGroup struct {
	sync.WaitGroup
}

func (r *WaitGroup) DoAction(cb func()) {
	r.Add(1)
	go func() {
		cb()
		r.Done()
	}()
}

func (r *WaitGroup) WaitActions() {
	r.Wait()
}