package tcpserver

import (
	"projectx/src/components/queue"
	. "projectx/src/config"
	"net"
	"sync"
	"projectx/src/util"
	"projectx/src/protocol"
)

func NewServer(q queue.MessageQueue) *server {
	return &server{
		q:q,
		cuid:0,
		clients:make(map[uint32]net.Conn),
	}
}

type server struct {
	q queue.MessageQueue
	listener net.Listener
	cl 	sync.RWMutex
	clients map[uint32]net.Conn
	cuid 	uint32
	cb 	*ClientEventCallback
}

func (r *server) Start(c TcpServerConfig, cb *ClientEventCallback) bool {
	l, err := net.Listen("tcp", c.Host)
	if err != nil {
		return false
	}
	r.listener = l

	accept := func() {
		for {
			clientConn, err := l.Accept()
			if err != nil {
				continue
			}
			go r.newconn(clientConn)
		}
	}

	go accept()
	return true
}

func (r *server) Stop() bool {
	return false
}

func (r *server) newconn(conn net.Conn) {
	r.cl.Lock()
	r.cuid = r.cuid + 1
	r.clients[r.cuid] = conn
	r.cl.Unlock()
	r.handlein(r.cuid, conn)
}

func (r *server) handlein(id uint32, conn net.Conn) {
	defer conn.Close()
	for {
		msg, router, err := util.ReadPacket(conn)
		if err != nil {

		}
		r.q.Push(id, id, msg)
	}
}

func (r *server) SendClient(id int32, msg *protocol.Message) {

}

func (r *server) BcClients(ids []int32, msg *protocol.Message) {

}

func (r *server) CloseClient(id int32, msg *protocol.Message) {

}

func (r *server) ConfigureClient(id int32, cfg *ClientConnectionConfigure) {

}
