package queue

import (
	"sync"
	"fmt"
	"projectx/src/protocol"
)

type MultiChans struct {
	chslock sync.Mutex
	chs 	map[uint32]QueueChanel
}

func NewMultiChans() *MultiChans{
	return &MultiChans{
		chs: make(map[uint32]QueueChanel),
	}
}

func (r *MultiChans) CreateChs(id uint32, size uint32) {
	r.chslock.Lock()
	r.chs[id] = QueueChanel{Ch:make(chan *protocol.Message, size)}
	r.chslock.Unlock()
}

func (r *MultiChans) ReleaseChs(id uint32) {
	r.chslock.Lock()
	delete(r.chs, id)
	r.chslock.Unlock()
}

func (r *MultiChans) Push(id uint32, dst int32, msg *protocol.Message) {

	r.chslock.Lock()
	if _, ok := r.chs[dst]; ok != true {
		fmt.Println("destination error : ", dst)
		r.chslock.Unlock()
		return
	}
	ch := r.chs[dst].Ch
	r.chslock.Unlock()

	ch <- msg;
}
