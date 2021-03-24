package main

import (
	"fmt"
	"testing"
)

func TestPacket(t *testing.T) {
	/*
		packet := &RoomListPacket{
			PacketType: 22,
			CurrentPage: 1,
			PerPage: 10,
		}

		b := convertPacketRoomList(packet)
		c := convertRoomList(&b)
		fmt.Println(packet)
		fmt.Println(c)
		if packet.PacketType != c.PacketType {
			fmt.Println("error")
		}

	*/

	// 비즈니스 로직
	// packet(비즈니스 처리한 패킷) -> RoomListPacket
	// bPacket -> PacketBuffer

	roomPacket := &RoomListPacket{
		CurrentPage: -1,
		PerPage:     1,
	}

	packet := &WritePacket{}
	packet.put(22, roomPacket)
	body := packet.merge()

	fmt.Println(body)

	readPacket := &ReadPacket{}
	readPacket.parse(body)

	roomPacket2 := &RoomListPacket{}
	roomPacket2.unMarshal(readPacket.PacketData)
	fmt.Printf("result %#v \n", roomPacket2)
	// -> send

	// --------------------- network ------------------------------------

	// --------------------- server side -----------------------------
	// receive
	//	queue := NewQueue()
	//	queue.Enqueue(b)

	// ------------------------------- logic process ----------------
}
