package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"unsafe"
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

type WritePacket struct {
	PacketType int32
	PacketData interface{}
}

type ReadPacket struct {
	PacketType int32
	PacketData []byte
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

/*
1. packet 보내기
- packet type 을 추가하여 전송
   -> packet type, data를 받아서 하나의 packet으로 만들어주기
   -> packet type 4byte로 정의해서 변경
   -> packet data -> json marshal
   -> type, data 연결해서 하나로 전송

2. packet 받기
- packet type, data로 분리하여 해당 type에 맞춰서 알맞은 로직으로 전달
   -> packet convert 할 때 4byte는 type으로 보고 string으로 변경해서 type 확인
   -> 나머지 byte 자르고 json unmarshal
*/

func (p *WritePacket) put(packetType int32, packetData interface{}) {
	p.PacketType = packetType
	p.PacketData = packetData
}

func (p *WritePacket) merge() []byte {
	var buf bytes.Buffer
	buf.Write(intToByte(p.PacketType))
	data, err := json.Marshal(p.PacketData)
	if err != nil {
		fmt.Println("merge marshal error: ", err)
	}
	buf.Write(data)

	return buf.Bytes()
}

func (p *ReadPacket) parse(packet Packet) error {
	buf := bytes.NewBuffer(packet)
	p.readType(buf)
	p.readData(buf)

	return nil
}

func (p *ReadPacket) readType(buffer *bytes.Buffer) error {
	var typeSize = unsafe.Sizeof(int32(1))
	buf := buffer.Bytes()
	packetTypeByte := buf[:typeSize]
	p.PacketType = int32(binary.BigEndian.Uint32(packetTypeByte))

	return nil
}

func (p *ReadPacket) readData(buffer *bytes.Buffer) error {
	var typeSize = unsafe.Sizeof(int32(1))
	buf := buffer.Bytes()
	p.PacketData = buf[typeSize:]

	return nil
}

func (p *RequestRoomListPacket) dataUnmarshal(data []byte) error {
	err := json.Unmarshal(data, p)
	if err != nil {
		return err
	}

	return nil
}

func intToByte(f int32) []byte {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, f)
	if err != nil {
		fmt.Println("binary.Write failed: ", err)
	}

	return buf.Bytes()
}

func float64ToByte(f float64) []byte {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, f)
	if err != nil {
		fmt.Println("binary.Write failed :", err)
	}

	return buf.Bytes()
}
