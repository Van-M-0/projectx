package queue

type ChannelMessage struct {
	Dst 	uint8 		//endpoint
	Id	uint32		//source id
	Data 	[]byte		//message data
}

type QueueChanel struct {
	Ch 	chan *ChannelMessage
}

type MessageQueue interface {
	Push(from uint32, data []byte)
}

