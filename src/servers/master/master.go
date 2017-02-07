package master

import (
	"projectx/src/components/queue"
	"projectx/src/config"
	"projectx/src/components/tcpserver"
	"fmt"
)

type Master interface {
	Start()
	Stop()
}

type master struct {
	q *queue.QueueRouter
	c config.MasterConfig
	server tcpserver.Server
}

func NewMaster() *master {
	master := new(master)
	master.c.Name = "master"
	master.c.Host = "127.0.1.0:9910"
	master.q = queue.NewMessageRouter()
	master.server = tcpserver.NewServer(master.q)
	return master
}

func (r *master) Start() {
	r.makerouter()
	r.server.Start(r.c.TcpServerConfig)
	r.run()
}

func (r *master) Stop() {
	r.server.Stop()
}

func (r *master) run() {
	for {
		r.handlecmd()
		r.handleio()
		r.handletimer()
	}
	fmt.Println("master run finish")
}

func (r *master) makerouter() {

}

func (r *master) handleio() {

}

func (r *master) handletimer() {

}

func (r *master) handlecmd() {

}
