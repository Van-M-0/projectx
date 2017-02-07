package util

import (
	"encoding/binary"
	"io"
	"net"
	"fmt"
	"bufio"
	"bytes"
	"projectx/src/protocol"
	"errors"
	"github.com/golang/protobuf/proto"
)

var (
	packetsizeerror = errors.New("packet size too small")
	packetdataerror = errors.New("packet data error")
)


func BinaryReadSize(reader io.Reader, buf []byte, size uint32) {

}



func test_parse_binary() {
	l, err := net.Listen("tcp", "127.0.0.1:4321")
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

		breader := bytes.NewReader(buf)
		var p person
		fmt.Println("person size : ", binary.Size(p))
		binary.Read(breader, binary.LittleEndian, &p)
		fmt.Println("person is : ", p.Len, string(p.Info[0:p.Len]))
	}

}

func ReadPacket(conn net.Conn, pb proto.Message) ([]byte, error) {
	hbuf := make([]byte, protocol.CommonHeaderLen)
	_, err := io.ReadFull(conn, hbuf)
	if err != nil {
		return nil, err
	}

	pbheader := protocol.CommonHeader{}
	reader := bytes.NewReader(hbuf)
	binary.Read(reader, binary.LittleEndian, &pbheader.Len)

	datasize := pbheader.Len - uint16(protocol.CommonHeaderLen)
	if datasize < 0 {
		fmt.Println("read data size error", datasize)
		return nil, packetsizeerror
	}

	data := make([]byte, datasize)
	if _, err := io.ReadFull(conn, data); err != nil {
		fmt.Println("read data error")
		return nil, packetdataerror
	}

	return data, nil
}
