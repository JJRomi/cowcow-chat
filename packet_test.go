package main

import (
	"fmt"
	"math/rand"
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

	maxSize := 10000 * 100
	var roomPacketData []RequestRoomListPacket = make([]RequestRoomListPacket, maxSize)
	for i := 0; i < maxSize; i++ {
		roomPacketData[i] = makeTestData(i)
	}

	writePacket := &WritePacket{}
	for i := 0; i < maxSize; i++ {
		go sendPacket(writePacket, roomPacketData[i])
	}
	fmt.Println("count : ", count)

	/**/

}

func makeTestData(id int) RequestRoomListPacket {
	return RequestRoomListPacket{
		Id:          id,
		CurrentPage: rand.Intn(100),
		PerPage:     rand.Intn(100),
	}
}

func sendPacket(packet *WritePacket, data interface{}) {
	packet.put(1, data)
	body := packet.merge()
	run(body)
}
