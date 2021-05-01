package main

import (
	"encoding/json"
	"fmt"
)

type Packet []byte

type RequestRoomListPacket struct {
	Id          int
	CurrentPage int
	PerPage     int
}

type RoomPacket struct {
	val int32
}

func convertPacketRoomList(p *RequestRoomListPacket) Packet {
	data, err := json.Marshal(p)
	if err != nil {
		fmt.Errorf("convertPacket error: %s", err)
	}

	return data
}

func convertRoomList(p *Packet) *RequestRoomListPacket {
	var room = &RequestRoomListPacket{}
	err := json.Unmarshal(*p, room)
	if err != nil {
		fmt.Errorf("convertRoomList error: %s", err)
	}

	return room
}
