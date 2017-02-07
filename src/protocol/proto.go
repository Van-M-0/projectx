package protocol

import (
	"projectx/src/protocol/baseproto"
	"github.com/golang/protobuf/proto"
)

// reference package, init it
var _ = baseproto.Null{}



type Message struct {
	proto.Message
}

func (m *Message) Name() string {
	return proto.MessageName(m.Message)
}



func GetProtoMessage(id int32) *Message {
	msg, ok := proto_map_reflection[id]
	if ok {
		return msg
	}
	return nil
}

func create_message(pb proto.Message) *Message {
	return &Message{Message:pb}
}



const (
	COMMON_HEADER 	= 1

	SERVER_REGISTER = 10001
	SERVER_GETALL_SERVERS = 10002
)


var (
	proto_map_reflection = map[int32]*Message {

		COMMON_HEADER: create_message(&baseproto.CommonHeader{}),

		SERVER_REGISTER: create_message(&baseproto.RegisterServer{}),
		SERVER_GETALL_SERVERS: create_message(&baseproto.AllServerInfo{}),
	}
)
