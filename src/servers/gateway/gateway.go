package gateway

import (
	"projectx/src/components/queue"
	"projectx/src/config"
	"projectx/src/components/tcpserver"
	"fmt"
	"net"
	"projectx/src/util"
	"projectx/src/protocol/baseproto"
	"projectx/src/protocol"
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
	backend   chan *protocol.Message
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

			err = util.WritePacket(conn, &baseproto.RegisterServer {

			}, 0)

			msg, _, err := util.ReadPacketByName(conn, "baseproto.AllServerInfo")

			for _, v := range (msg.(&baseproto.AllServerInfo).Servers) {
				r.serinfos.add(&serverpeer{
					host: v.Ip,
					port: v.Port,
					server_type: v.Type,
					id: v.Id,
				})
			}

			go r.readservermsg(conn, util.SERVER_TYPE_MASTER, util.SERVER_TYPE_MASTER)
		})
	})
	start_wait.WaitActions()

	// lobby
	start_wait.DoAction(func() {
		connect(":9399", func(conn net.Conn, err error) {
			go r.readservermsg(conn, util.SERVER_TYPE_LOBBY, util.SERVER_TYPE_LOBBY)
			r.registerclientrouter(util.SERVER_TYPE_LOBBY, 1024, conn)
		})
	})

	// game servers
	for _, server := range(r.serinfos.servers) {
		if server.server_type == util.SERVER_TYPE_GAMESERVER {
			start_wait.DoAction(func() {
				connect(server.host, func(conn net.Conn, err error) {

					s := r.serinfos.getbyid(server.id)
					r.serinfos.setconn(s.id, conn)
					r.registerclientrouter(s.id, 4096, conn)

					go r.readservermsg(conn, util.SERVER_TYPE_GAMESERVER, s.id)
				})
			})

		}
	}
	start_wait.WaitActions()

	r.chs.CreateChs(util.SERVER_TYPE_GATEWAY, 1024)
}

func (r *gateway) registerclientrouter(id int32, size int32, conn net.Conn) {
	ch := r.chs.CreateChs(id, size)
	go func() {
		select {
		case msg := <- ch:
			util.WritePacket(conn, msg, 1)
		}
	}()
}

func (r *gateway) readservermsg(conn net.Conn, t int32, id int32) {

	go r.handleio()

	for {
		msg, _, err := util.ReadPacket(conn)
		if err != nil {

		}
		r.backend <- msg
	}
}

func (r *gateway) handleio() {
	for {
		select {
		}
	}
}
