package game

import (
	"projectx/src/components/queue"
	"projectx/src/config"
	"projectx/src/components/tcpserver"
	"fmt"
	"net"
	"projectx/src/util"
	"projectx/src/protocol"
)

type game struct {
	chs       *queue.MultiChans
	c         config.MasterConfig
	server    tcpserver.Server
	backend   chan *protocol.Message
}

func NewGame() *game {
	gw := new(game)
	gw.c.Name = "master"
	gw.c.Host = "127.0.1.0:9910"
	gw.chs = queue.NewMultiChans()
	gw.server = tcpserver.NewServer(gw.chs)
	return gw
}

func (r *game) Start() {
	r.server.Start(r.c.TcpServerConfig, &tcpserver.ClientEventCallback{
		OnConnect: func(id int32, conn net.Conn) { r.OnConnect(id, conn)},
		OnClose: func(id int32) {r.OnClose(id)},
		OnError: func(id int32) {r.OnError(id)},
	})
	r.run()
}

func (r *game) Stop() {
	r.server.Stop()
}

func (r *game) run() {
	fmt.Println("master run finish")
}

func (r *game) OnConnect(connid int32, conn net.Conn) {
	msg, _, err := util.ReadPacketByName(conn, "baseproto.RegisterServer")
	if err != nil {

	}

}

func (r *game) OnClose(id int32) {

}

func (r *game) OnError(id int32) {

}
