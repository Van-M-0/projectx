package queue

import "projectx/src/protocol"

type QueueChanel struct {
	Ch 	chan *protocol.Message
}

type MessageQueue interface {
	Push(from uint32, dst int32, msg *protocol.Message)
}

