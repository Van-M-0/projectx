package queue

import "projectx/src/protocol"

type QueueChanel struct {
	Ch 	chan *protocol.Message
}

type MessageQueue interface {
	Push(from uint32, dst int32, msg *protocol.Message)
}

type ChannlPoint struct {
	Id 	interface{}
	Type 	int32
}

type ChannelMessage struct {
	Src 	ChannlPoint
	Dest 	ChannlPoint
	Ch 	chan *protocol.Message
}

type Channel interface {
	Push(*ChannelMessage)
	Pull() (*ChannelMessage, error)
}

