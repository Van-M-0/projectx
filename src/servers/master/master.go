package master

import (
	"projectx/src/components/queue"
	"projectx/src/config"
	"projectx/src/components/tcpserver"
	"fmt"
	"net"
	"projectx/src/util"
	"projectx/src/protocol"
)

type GateWay interface {
	Start()
	Stop()
}

type master struct {
	chs       *queue.MultiChans
	c         config.MasterConfig
	server    tcpserver.Server
	backend   chan *protocol.Message
}

func NewGateWay() *master {
	gw := new(master)
	gw.c.Name = "master"
	gw.c.Host = "127.0.1.0:9910"
	gw.chs = queue.NewMultiChans()
	gw.server = tcpserver.NewServer(gw.chs)
	return gw
}

func (r *master) Start() {
	r.connectserver()
	r.server.Start(r.c.TcpServerConfig, &tcpserver.ClientEventCallback{
		OnConnect: func(id int32) { r.OnConnect(id)},
		OnClose: func(id int32) {r.OnClose(id)},
		OnError: func(id int32) {r.OnError(id)},
	})
	r.run()
}

func (r *master) Stop() {
	r.server.Stop()
}

func (r *master) run() {
	fmt.Println("master run finish")
}

func (r *master) registerclientrouter(id int32, size int32, serconn net.Conn) {
	ch := r.chs.CreateChs(id, size)
	go func() {
		select {
		case msg := <- ch:
			util.WritePacket(serconn, msg, 1)
		}
	}()
}

func (r *master) OnConnect(id int32) {

}

func (r *master) OnClose(id int32) {

}

func (r *master) OnError(id int32) {

}

func (r *master) handelbackend() {
	for {
		select {
		case msg := <- r.backend:
			msg
		}
	}
}
