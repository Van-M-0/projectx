package protocol

import (
	"encoding/binary"
)

const (
	MASTER 	= 1
	GATEWAY = 2
	GAME	= 3
	CACHE	= 4
	DBFE	= 5
)

var (
	CommonHeaderLen = binary.Size(CommonHeader{})
	//ClientHeaderLen = binary.Size(ClientHeader{})
	//ServerHeaderLen = binary.Size(ServerHeader{})
)

///*
//	CommonHeader +
//*/
//type CommonHeader struct {
//	Len 	uint16
//}
//
//type RouterHeader struct {
//	Dst 	uint8
//}
//
//type ClientHeader struct {
//	Len 	uint16
//	ver 	uint8
//	magic 	uint8
//}
//
//type ServerHeader struct {
//	Len 	uint16
//}

//func EncodeHeader (d []byte) []byte {
//	h := ClientHeader{}
//	h.magic = 1
//	h.ver = 0
//	h.Len = len(h) + len(d)
//	buf := new(bytes.Buffer)
//	binary.Write(buf, binary.LittleEndian, h)
//	rb := append([]byte{}, buf.Bytes())
//	rb = append(rb, d)
//	return rb
//}
//
//func DecodeHeader (d []byte) ([]byte, uint16) {
//	return d, 10
//}


