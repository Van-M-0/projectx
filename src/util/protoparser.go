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
	"projectx/src/protocol/baseproto"
)

var (
	packetsizeerror = errors.New("packet size too small")
	packettypeerror = errors.New("packet proto type err")
	packetdataerror = errors.New("packet data error")

	common_header_size = proto.Size(baseproto.CommonHeader{})
)

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

func ReadPacket(conn net.Conn) (*protocol.Message, int32, error) {
	buf := make([]byte, common_header_size)
	_, err := io.ReadFull(conn, buf)
	if err != nil {
		return nil, 0, err
	}

	var header baseproto.CommonHeader
	header = protocol.GetProtoMessage(protocol.COMMON_HEADER)
	err = proto.Unmarshal(buf, header)
	if err != nil {
		return nil, 0, err
	}

	msg := protocol.GetProtoMessage(header.Id)
	if msg == nil {
		return nil, 0, packettypeerror
	}

	data_size := header.Len - common_header_size
	if data_size < 0 {
		return nil, 0, packetsizeerror
	}

	data := make([]byte, data_size)
	_, err = io.ReadFull(conn, data)
	if err != nil {
		return nil, 0, err
	}

	err = proto.Unmarshal(data, msg)
	if err != nil {
		return nil, 0, err
	}

	return msg, header.Router, nil
}

func WritePacket(conn net.Conn, msg *protocol.Message, flag uint32) (error) {
	var header baseproto.CommonHeader
	header.Len = common_header_size
	header, err := proto.Marshal(header)
	if err != nil {
		return err
	}

	body, err := proto.Marshal(msg.Message)
	if err != nil {
		return err
	}

	packet := append([]byte{}, header)
	packet = append(packet, body)

	conn.Write(packet)

	return nil
}
