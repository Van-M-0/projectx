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

func (r *MultiChans) CreateChs(id uint32, size uint32) chan *protocol.Message{
	r.chslock.Lock()
	defer r.chslock.Unlock()
	r.chs[id] = QueueChanel{Ch:make(chan *protocol.Message, size)}
	return r.chs[id].Ch
}

func (r *MultiChans) ReleaseChs(id uint32) {
	r.chslock.Lock()
	defer r.chslock.Unlock()
	delete(r.chs, id)
}

func (r *MultiChans) Push(id uint32, dst int32, msg *protocol.Message) {

	r.chslock.Lock()
	defer r.chslock.Unlock()

	if _, ok := r.chs[dst]; ok != true {
		fmt.Println("destination error : ", dst)
		r.chslock.Unlock()
		return
	}
	ch := r.chs[dst].Ch
	ch <- msg;
}
