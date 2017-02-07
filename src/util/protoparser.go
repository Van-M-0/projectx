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

	common_header_size = binary.Size(protocol.CommonHeader{})
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

func ReadPacket(conn net.Conn, pb proto.Message) (error) {
	buf := make([]byte, common_header_size)
	_, err := io.ReadFull(conn, buf)
	if err != nil {
		return nil, err
	}

	var header protocol.CommonHeader
	err = proto.Unmarshal(buf, header)
	if err != nil {
		return err
	}

	data_size := header.Len - common_header_size
	if data_size < 0 {
		return packetsizeerror
	}

	data := make([]byte, data_size)
	_, err = io.ReadFull(conn, data)
	if err != nil {
		return err
	}
	err = proto.Unmarshal(data, pb)
	if err != nil {
		return err
	}

	return nil
}

func WritePacket(conn net.Conn, pb proto.Message, flag uint32) ([]byte, error) {
	var pbheader protocol.CommonHeader
	pbheader.Len = common_header_size
	header, err := proto.Marshal(pbheader)
	if err != nil {
		return nil, err
	}

	body, err := proto.Marshal(pb)
	if err != nil {
		return nil, err
	}

	packet := append([]byte{}, header)
	packet = append(packet, body)

	conn.Write(packet)

	return packet, nil
}
