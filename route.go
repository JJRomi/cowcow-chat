package main

import "fmt"

/*
type 을 받고
type 에 따라서 비즈니스 로직 타게 하는 역할
*/

// networkCard(packet) -> os -> golang([]byte) -> route -> logic
func run(packet []byte) error {
	readPacket := &ReadPacket{}
	readPacket.parse(packet)

	// 조건문..
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

	err := roomListPacket.unMarshal(packetData)
	if err != nil {
		return err
	}

	fmt.Printf("result %+v \n", roomListPacket)

	return nil
}

// requestRoomList
// successRoomList
// failRoomList code, msg
