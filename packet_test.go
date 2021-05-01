package main

import (
	"fmt"
	"testing"
)

func TestPacket(t *testing.T) {
	packet := &RequestRoomListPacket{
		Id:          1,
		CurrentPage: 1,
		PerPage:     10,
	}

	packetRoomList := convertPacketRoomList(packet)
	roomList := convertRoomList(&packetRoomList)
	fmt.Println(packet)
	fmt.Println(roomList)
	if packet.Id != roomList.Id {
		fmt.Println("RoomListPacket error")
	}
}
