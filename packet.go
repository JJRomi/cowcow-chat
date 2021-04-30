package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"unsafe"
)

// header / data
// type | length
// type | full length | string_len | string_data | int_len | int_data
// client -> object -> convert byte[] -> server(byte[] -> object) -> object -> logic

type Packet []byte

//type Packet struct {
//	data []byte
//}

// type 0
// 1. 4(type) | 4(cur_page) | cur_page | 4(perPage) | perPage
// 2. 8(type)  | 4(full data length) | 4(cur_page) | cur_page | 4(perPage) | perPage
type RequestRoomListPacket struct {
	// PacketType  int
	Id          int
	CurrentPage int
	PerPage     int
}

type SuccessRoomListPacket struct {
}

type FailRoomListPacket struct {
}

type RoomPacket struct {
	val int32
}

type CreateRoomPacket struct {
}

type WritePacket struct {
	PacketType int32
	PacketData interface{} // <T>
}

type ReadPacket struct {
	PacketType int32
	PacketData []byte // <T>
}

/*
func (r *test) put(data interface{}) {
	r.data = data
}
*/

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
func PacketParser(p *Packet) *PacketBuffer {
	return nil
}
*/
/*
1. packet 보내기
- packet type 을 추가하여 전송
	-> packet type 하고 data 를 받아서 하나의 packet 으로 만들어주기
	-> packet type 4byte 로 정의해서 변경해주고
	-> packet data -> json marshal 해주기
	-> type 하고 data 를 연결해서 하나로 전송

2. packet 받기
- packet type 하고 data 로 분리하여 해당 type 에 맞춰서 알맞는 로직으로 보내기
	-> packet convert 해줄때 앞에 4byte 는 type 으로 보고 string 으로 변경해서 type 확인
	-> 나머지는 byte 자르고 json unmarshal
*/

func (p *WritePacket) put(packetType int32, packetData interface{}) {
	p.PacketType = packetType
	p.PacketData = packetData
}

func (p *WritePacket) merge() []byte {
	// 타임 사이즈 구하기
	// append 처리
	// size(byte) + packetType(byte) + packet
	var buf bytes.Buffer
	//var typeSize int = unsafe.Sizeof(int(0))
	//var typeSize int = 1
	// size := unsafe.Sizeof(typeSize)
	//size := unsafe.Sizeof(int(1))
	//fmt.Println(size)
	buf.Write(intToByte(p.PacketType))
	data, err := json.Marshal(p.PacketData)
	if err != nil {
		fmt.Println(err)
	}
	buf.Write(data)

	return buf.Bytes()

	/*var num1 int8  = 1
	var num2 int16 = 1
	var num3 int32 = 1
	var num4 int64 = 1

	fmt.Println(unsafe.Sizeof(num1)) // 1
	fmt.Println(unsafe.Sizeof(num2)) // 2
	fmt.Println(unsafe.Sizeof(num3)) // 4
	fmt.Println(unsafe.Sizeof(num4)) // 8


	*/
}

func (p *ReadPacket) parse(packet Packet) error {
	// 분리
	// type(string) , data([]byte}
	buf := bytes.NewBuffer(packet)
	p.readType(buf)
	p.readData(buf)
	return nil
}

//buf := bytes.NewBuffer(b) // b is []byte
//myfirstint, err := binary.ReadVarint(buf)
//anotherint, err := binary.ReadVarint(buf)
func (p *ReadPacket) readType(buffer *bytes.Buffer) error {
	var typeSize = unsafe.Sizeof(int32(1))
	// buffer를 typesize만큼 자르는 부분
	buf := buffer.Bytes()
	//	fmt.Println(typeSize)
	packetTypeByte := buf[:typeSize]
	//	fmt.Println("packet: ", packetTypeByte)
	p.PacketType = int32(binary.BigEndian.Uint32(packetTypeByte))
	//	fmt.Println("packet int: ", p.PacketType)
	return nil
}

func (p *ReadPacket) readData(buffer *bytes.Buffer) error {
	var typeSize = unsafe.Sizeof(int32(1))
	buf := buffer.Bytes()
	p.PacketData = buf[typeSize:]
	return nil
}

func (p *RequestRoomListPacket) unMarshal(data []byte) error {
	err := json.Unmarshal(data, p)
	if err != nil {
		return err
	}
	return nil
}
