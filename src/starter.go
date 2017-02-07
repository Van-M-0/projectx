package main

import (
	"fmt"
	"projectx/src/protocol"
)


func main() {
	fmt.Println("main run !!")
	fmt.Println("main stop")

	msg := protocol.GetProtoMessage(protocol.COMMON_HEADER)
	fmt.Println("name ", msg.Name())
}

