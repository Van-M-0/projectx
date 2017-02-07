package gateway

import (
	"projectx/src/components/queue"
	"projectx/src/config"
	"projectx/src/components/tcpserver"
	"fmt"
	"net"
	"projectx/src/protocol"
	"projectx/src/util"
	"projectx/src/protocol/baseproto"
)

type GateWay interface {
	Start()
	Stop()
}

type gateway struct {
	chs       *queue.MultiChans
	c         config.MasterConfig
	server    tcpserver.Server
	serinfos  serverlist
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
			err = util.WritePacket(conn, &baseproto.RegisterServer {

			}, 0)

			// get all server - lists
			msg, router, err := util.ReadPacketByName(conn, "baseproto.AllServerInfo")

			for _, v := range (msg.(&baseproto.AllServerInfo).Servers) {
				r.serinfos.add(&serverpeer{
					host: v.Ip,
					port: v.Port,
					server_type: v.Type,
					id: v.Id,
				})
			}
			s := r.serinfos.getbytype(util.SERVER_TYPE_MASTER)
			r.serinfos.setconn(s.id, conn)
			r.chs.CreateChs(s.id, 1024)
		})
	})
	start_wait.WaitActions()

	// lobby
	start_wait.DoAction(func() {
		connect(":9399", func(conn net.Conn, err error) {
			s := r.serinfos.getbytype(util.SERVER_TYPE_LOBBY)
			r.serinfos.setconn(s.id, conn)
			r.chs.CreateChs(s.id, 1024)
		})
	})

	// game servers
	for _, server := range(r.serinfos.servers) {
		if server.server_type == util.SERVER_TYPE_GAMESERVER {
			start_wait.DoAction(func() {
				connect(server.host, func(conn net.Conn, err error) {
					s := r.serinfos.getbyid(server.id)
					r.serinfos.setconn(s.id, conn)
					r.chs.CreateChs(s.id, 1024)
				})
			})
		}
	}
	start_wait.WaitActions()

	r.chs.CreateChs(-1, 4096)
}


func (r *gateway) handleio() {

}
