package gateway

import (
	"sync"
	"net"
)

type serverpeer struct {
	server_type 	int32
	id 		int32
	host 		string
	port 		int32
	conn 		net.Conn
}

type serverlist struct {
	sync.Mutex
	servers 	map[int32]*serverpeer
}

func (r *serverlist) add(s *serverpeer) {
	r.Lock()
	r.servers[s.id]	= s
	r.Unlock()
}

func (r *serverlist) remove(id int32) {
	r.Lock()
	delete(r.servers, id)
	r.Unlock()
}

func (r *serverlist) getbyid(id int32) *serverpeer {
	r.Lock()
	var s *serverpeer
	if _, ok := r.servers[id]; ok == true {
		s = r.servers[id]
	}
	r.Unlock()
	return s
}

func (r *serverlist) getbytype(t int32) *serverpeer {
	r.Lock()
	var s *serverpeer
	for _, v := range r.servers {
		if v.server_type == t {
			s = v
			break
		}
	}
	r.Unlock()
	return s
}

func (r *serverlist) setconn(id int32, conn net.Conn) {
	s := r.getbyid(id)
	s.conn = conn
}
