package tcpserver

import "projectx/src/config"

const (
	NET_EVENT_NEWCONN = 1
	NET_EVENT_CLSCONN = 2
	NET_EVENT_READ	  = 3
	NET_EVENT_WRITE   = 4
)

type Client interface {

}

type Server interface {
	Start(config.TcpServerConfig) bool
	Stop() bool
}
