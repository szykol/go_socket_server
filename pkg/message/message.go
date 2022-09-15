package message

import (
	"bytes"
	"encoding/gob"
)

const (
	MSG_T_EXAMPLE          uint16 = 0
	TYPE_HEADER_SIZE_BYTES        = 2
)

func DecodeValue[T any](buf []byte) (T, error) {
	var parsedValue T
	r := bytes.NewBuffer(buf)
	decoder := gob.NewDecoder(r)
	err := decoder.Decode(&parsedValue)
	return parsedValue, err
}

type Payload interface {
	ToBytes() ([]byte, error)
	FromBytes([]byte) error
	SizeInBytes() int
}

type VectorPayload struct {
	X int16
	Y int16
}

func (p *VectorPayload) ToBytes() ([]byte, error) {
	w := new(bytes.Buffer)
	encoder := gob.NewEncoder(w)
	err := encoder.Encode(p.X)
	if err != nil {
		return nil, err
	}
	err = encoder.Encode(p.Y)
	if err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

func (p *VectorPayload) FromBytes(buf []byte) error {
	r := bytes.NewBuffer(buf)
	decoder := gob.NewDecoder(r)
	err := decoder.Decode(&p.X)
	if err != nil {
		return err
	}
	return decoder.Decode(&p.Y)
}

func (p *VectorPayload) SizeInBytes() int {
	return 8
}

type Message struct {
	Type    uint16
	Payload Payload
}

func (m *Message) ToBytes() ([]byte, error) {
	w := new(bytes.Buffer)
	encoder := gob.NewEncoder(w)
	err := encoder.Encode(m.Type)
	if err != nil {
		return nil, err
	}
	payloadBytes, _ := m.Payload.ToBytes()
	messageBytes := append(w.Bytes()[:TYPE_HEADER_SIZE_BYTES], payloadBytes...)
	return messageBytes, nil
}
