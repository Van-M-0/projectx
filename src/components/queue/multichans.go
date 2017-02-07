package queue

import (
	"sync"
	"fmt"
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
	r.chs[id] = QueueChanel{Ch:make(chan *ChannelMessage, size)}
	r.chslock.Unlock()
}

func (r *MultiChans) ReleaseChs(id uint32) {
	r.chslock.Lock()
	delete(r.chs, id)
	r.chslock.Unlock()
}

func (r *MultiChans) Push(id uint32, data []byte) {
	m := &ChannelMessage{
		Dst:data[0],
		Id:id,
		Data:data[1:],
	}

	r.chslock.Lock()
	if _, ok := r.chs[m.Dst]; ok != true {
		fmt.Println("destination error : ", m.Dst)
		r.chslock.Unlock()
		return
	}
	var ch chan *ChannelMessage
	ch = r.chs[m.Dst].Ch
	r.chslock.Unlock()

	ch <- m;
}
