package gateway

import (
	"projectx/src/components/queue"
	"projectx/src/config"
	"projectx/src/components/tcpserver"
	"fmt"
	"net"
	"projectx/src/protocol"
	"projectx/src/util"
)

type GateWay interface {
	Start()
	Stop()
}

type gateway struct {
	chs *queue.MultiChans
	c config.MasterConfig
	server tcpserver.Server
	allserver *protocol.GetAllServerInfo
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
	r.connectserver()
	r.server.Start(r.c.TcpServerConfig)
	r.run()
}

func (r *gateway) Stop() {
	r.server.Stop()
}

func (r *gateway) run() {
	for {
		r.handleio()
	}
	fmt.Println("gateway run finish")
}

func (r *gateway) connectserver() {

	connect := func(dsthost string, cb func()) {
		conn, err := net.Dial("tcp", dsthost)
		if err != nil {
			cb(nil, err)
		}

		if cb != nil {
			cb(conn, nil)
		}
	}

	start_wait := new(util.WaitGroup)

	// master
	start_wait.DoAction(func() {
		connect(":9090", func(conn net.Conn, err error) {
			// register self server info
			err = util.WritePacket(conn, &protocol.ServerInfo {

			}, 0)

			// get all server - lists
			err = util.ReadPacket(conn, r.allserver)
		})
	})
	start_wait.WaitActions()

	// lobby
	start_wait.DoAction(func() {
		connect(":9399", func(conn net.Conn, err error) {
		})
	})

	// game servers
	gameservers :=[]*protocol.ServerInfo{}
	for _, server := range(r.allserver.Servers) {
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

func (r *gateway) handleio() {

}
