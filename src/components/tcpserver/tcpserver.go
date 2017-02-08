package tcpserver

import (
	"projectx/src/config"
	"projectx/src/protocol"
)

type Client interface {

}

type ClientEventCallback struct {
	OnConnect	func(id int32)
	OnClose 	func(id int32)
	OnError		func(id int32)
}

type ClientConnectionConfigure struct {

}

type Server interface {
	Start(config.TcpServerConfig, ec *ClientEventCallback) bool
	Stop() bool

	SendClient(id int32, msg *protocol.Message)
	BcClients(ids []int32, msg *protocol.Message)
	CloseClient(id int32, msg *protocol.Message)
	ConfigureClient(id int32, cfg *ClientConnectionConfigure)
}
