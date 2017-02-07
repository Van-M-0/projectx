package queue

type QueueRouter struct {
	channels map[uint32]QueueChanel
}

func NewMessageRouter() *QueueRouter {
	return &QueueRouter{
		channels: make(map[uint32]QueueChanel),
	}
}

func (r *QueueRouter) AddRouter(path uint32) {
	r.channels[path] = QueueChanel{Ch:make(chan *ChannelMessage, 1024)}
}

func (r *QueueRouter) RemoveRouter(path uint32) {
	delete(r.channels, path)
}

func (r *QueueRouter) Push(src, dst, id uint32, data []byte) {
	//chanel := &r.channels[dst]
	//chanel.Ch <- &ChannelMessage{
	//	src:src,
	//	dst:dst,
	//	id:id,
	//	Data:data,
	//}
}
