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


func (r *gateway) connectlobby() {
	_, err := net.Dial("tcp", "192.168.1.1:9800")
	if err != nil {
		log.Fatal("connect lobby err ", err)
	}

}

func (r *gateway) connectgameserver() {

}

func (r *gateway) connectserver() {
	connect := func(host string) {
		conn, err := net.Dial("tcp", host)
		if err != nil {
			log.Fatal("connect server err, ", host)
		}
		r.readfromserver(conn)
	}

	go connect("192.168.1.10:9910")
}

func (r *gateway) readfromserver(conn net.Conn) {
	defer conn.Close()
	for {
		data, err := util.ReadPacket(conn)
		if err != nil {
			log.Fatal("read pacekt err ", err)
		}
		r.chs.Push(0, data)
	}
}

func (r *gateway) HandleRoutine(data []byte) {

}

func (r *gateway) handleio() {

}

func (r *gateway) handletimer() {

}

func (r *gateway) handlecmd() {

}
