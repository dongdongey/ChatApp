package client

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Data struct {
	nameSize    uint8
	name        []byte
	payloadSize uint32
	payload     []byte
}

func (d *Data) SetName(name []byte) {
	var nameSize = len(name)
	if nameSize > 255 {
		fmt.Println("Data is too big!")
		return
	}
	d.nameSize = uint8(len(name))
	d.name = name
}

func (d *Data) SetPayload(payload []byte) {
	var payloadSize = len(payload)
	if payloadSize > 4294967295 {
		fmt.Println("Data is too big!")
		return
	}
	d.payloadSize = uint32(len(payload))
	d.payload = payload
}

func (d *Data) ToBytes() []byte {
	// 버퍼 생성
	buf := new(bytes.Buffer)
	// 버퍼에 이름 크기 작성
	binary.Write(buf, binary.LittleEndian, d.nameSize)
	// 버퍼에 이름 작성
	buf.Write(d.name)
	// 버퍼에 내용 크기 작성
	binary.Write(buf, binary.LittleEndian, d.payloadSize)
	// 버퍼에 내용 작성
	buf.Write(d.payload)
	return buf.Bytes()
}

func FromBytes(data []byte) *Data {
	// 버퍼 Reader 생성
	buf := bytes.NewReader(data)

	// uint8 크기의 데이터 읽어옴
	var nameSize uint8
	binary.Read(buf, binary.LittleEndian, &nameSize)
	// 이름 크기만큼의 이름 버퍼 생성하고 읽어옴
	name := make([]byte, nameSize)
	buf.Read(name)

	// uint64 크기의 데이터 읽어옴
	var payloadSize uint32
	binary.Read(buf, binary.LittleEndian, &payloadSize)
	// 내용 크기만큼의 내용 버퍼 생성하고 읽어옴
	payload := make([]byte, payloadSize)
	buf.Read(payload)

	return &Data{nameSize: nameSize, name: name, payloadSize: payloadSize, payload: payload}
}
