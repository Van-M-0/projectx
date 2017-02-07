package gateway

import (
	"projectx/src/components/queue"
	"projectx/src/config"
	"projectx/src/components/tcpserver"
	"fmt"
	"net"
	"log"
	"projectx/src/protocol"
	"projectx/src/util"
	"github.com/golang/protobuf/proto"
)

type GateWay interface {
	Start()
	Stop()
}

type gateway struct {
	chs *queue.MultiChans
	c config.MasterConfig
	server tcpserver.Server
}

func NewGateWay() *gateway {
	gw := new(gateway)
	gw.c.Name = "master"
	gw.c.Host = "127.0.1.0:9910"
	gw.chs = queue.NewMultiChans()
	gw.server = tcpserver.NewServer(gw.chs)
	return gw
}

func (r *gateway) Start() {
	r.chs.CreateChs(protocol.GATEWAY, 1024)
	r.chs.CreateChs(protocol.MASTER, 1024)
	r.connectserver()
	r.server.Start(r.c.TcpServerConfig)
	r.run()
}

func (r *gateway) Stop() {
	r.server.Stop()
}

func (r *gateway) run() {
	for {
		r.handlecmd()
		r.handleio()
		r.handletimer()
	}
	fmt.Println("gateway run finish")
}

func (r *gateway) connectserver() {

	const selfinfo = &protocol.ServerInfo{

	}

	connect := func(dsthost string, cb func()) {
		conn, err := net.Dial("tcp", dsthost)
		if err != nil {
			cb(nil, err)
		}

		err = util.WritePacket(conn, selfinfo, 0)
		if err != nil {
			cb(nil, err)
		}

		if cb != nil {
			cb(conn, nil)
		}
	}

	start_wait := new(util.WaitGroup)
	var allservers protocol.GetAllServerInfo

	// master
	start_wait.DoAction(func() {
		connect(":9090", func(conn net.Conn, err error) {

		})
	})

	// lobby
	start_wait.DoAction(func() {
		connect(":9399", func(conn net.Conn, err error) {
			err = util.ReadPacket(conn, &allservers)
		})
	})

	start_wait.WaitActions()

	// game servers
	gameservers :=[]*protocol.ServerInfo{}
	for _, server := range(allservers.Servers) {
		if server.Type == "game server"	{
			gameservers = append(gameservers, server)
		}
	}

	if len(gameservers) <= 0 {
		return
	}

	for _, gs := range(gameservers) {
		start_wait.DoAction(func() {
			connect(gs.Ip, func(conn net.Conn, err error) {

			})
		})
	}

	start_wait.WaitActions()
}

func (r *gateway) HandleRoutine(data []byte) {

}

func (r *gateway) handleio() {

}

func (r *gateway) handletimer() {

}

func (r *gateway) handlecmd() {

}
