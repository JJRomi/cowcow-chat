package main

import (
	"math/rand"
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

	/*roomPacket := &RequestRoomListPacket{
		CurrentPage: -1,
		PerPage:     1,
	}*/
	/*maxSize := 10000 * 100
	var roomPacketData []RequestRoomListPacket = make([]RequestRoomListPacket, maxSize)
	for i := 0; i < maxSize; i++ {
		roomPacketData[i] = makeTestData(i)
	}

	packet := &WritePacket{}
	for i := 0; i < maxSize; i++ {
		go sendPacket(packet, roomPacketData[i])
	}
	fmt.Println("count: ", count)*/

	/*
		readPacket := &ReadPacket{}
		readPacket.parse(body)

		roomPacket2 := &RoomListPacket{}
		roomPacket2.unMarshal(readPacket.PacketData)
		fmt.Printf("result %#v \n", roomPacket2)
	*/

	// -> send

	// --------------------- network ---------------------------------

	// --------------------- server side -----------------------------
	//  receive
	//	queue := NewQueue()
	//	queue.Enqueue(b)

	// --------------------- logic process ---------------------------
}

func sendPacket(packet *WritePacket, data interface{}) {
	packet.put(1, data)
	//packet.put(1, roomPacketData[i])
	body := packet.merge()
	run(body)
}

func makeTestData(id int) RequestRoomListPacket {
	return RequestRoomListPacket{
		CurrentPage: rand.Intn(100),
		PerPage:     rand.Intn(100),
		Id:          id,
	}
}
