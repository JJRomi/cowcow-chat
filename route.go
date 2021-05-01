package main

import "fmt"

func run(packet []byte) error {
	readPacket := &ReadPacket{}
	readPacket.parse(packet)

	switch readPacket.PacketType {
	case 1:
		readRoomList(readPacket.PacketData)
		break
	default:
		break
	}

	return nil
}

var count int = 0

func readRoomList(packetData []byte) error {
	count++
	roomListPacket := &RequestRoomListPacket{}

	err := roomListPacket.dataUnmarshal(packetData)
	if err != nil {
		return err
	}

	fmt.Printf("read room list result %+v \n", roomListPacket)

	return nil
}
