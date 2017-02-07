package util

import (
	"net"
	"fmt"
	"bufio"
	"encoding/binary"
	"io"
	"github.com/golang/protobuf/proto"
	"projectx/src/protocol"
)

func Test_proto() {
	l, err := net.Listen("tcp", ":4321")
	if err != nil {
		fmt.Println("listen err ", err)
	}
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("accept err ", err)
	}

	fmt.Println("new conn incoming")

	ta := []byte("hello world")
	fmt.Println("hello world ", ta)

	type person struct {
		Len int32
		Info [32]byte
	}

	reader := bufio.NewReader(conn)
	for {
		var messagelen int32
		err := binary.Read(reader, binary.LittleEndian, &messagelen)
		if err != nil {
			continue
		}
		fmt.Println("message len : ", messagelen, reader.Buffered())

		buf := make([]byte, messagelen - 4)
		_, err = io.ReadFull(reader, buf)
		if err != nil {
			continue
		}

		fmt.Println("buf is ; ", buf)

		p := &protocol.Person{}
		err = proto.Unmarshal(buf, p)
		if err != nil {
			fmt.Println("unmarshal err", err)
		}
	}
}
