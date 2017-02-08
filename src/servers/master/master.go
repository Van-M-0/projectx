package master

import (
	"projectx/src/components/queue"
	"projectx/src/config"
	"projectx/src/components/tcpserver"
	"fmt"
	"net"
	"projectx/src/util"
	"projectx/src/protocol"
	"projectx/src/protocol/baseproto"
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
	r.server.Start(r.c.TcpServerConfig, &tcpserver.ClientEventCallback{
		OnConnect: func(id int32, conn net.Conn) { r.OnConnect(id, conn)},
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

func (r *master) updateserverinfo(msg *baseproto.RegisterServer, add bool) {

}

func (r *master) OnConnect(connid int32, conn net.Conn) {
	msg, router, err := util.ReadPacketByName(conn, "baseproto.RegisterServer")
	if err != nil {

	}

	server := msg.(&baseproto.RegisterServer{}).Server
	r.updateserverinfo(server, true)

	ch := r.chs.CreateChs(server.Id, 1024)

	if server.Type == util.SERVER_TYPE_LOBBY {

	}
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
