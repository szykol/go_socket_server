package receiver

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"szykol/sockets/pkg/message"
)

type Receiver interface {
	ReadBytes() ([]byte, error)
	ReadString() (string, error)
}

type BufferReader interface {
	ReadBytes(byte) ([]byte, error)
	ReadString(byte) (string, error)
}

type SocketReceiver struct {
	reader BufferReader
	Delim  byte
}

func (s *SocketReceiver) ReadBytes() ([]byte, error) {
	return s.reader.ReadBytes(s.Delim)
}

func (s *SocketReceiver) ReadMessage(connection io.Reader) (*message.Message, error) {
	typeValue, err := s.readType(connection)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Wrong response when reading type bytes: %v", err))
	}

	var payload message.Payload
	if typeValue == message.MSG_T_EXAMPLE {
		payload = &message.VectorPayload{}
	} else {
		return nil, errors.New(fmt.Sprintf("Wrong type value: %v", typeValue))
	}

	readPayloadBuffer, _ := s.readNBytes(connection, payload.SizeInBytes())
	// if err != nil {
	// 	return nil, errors.New(fmt.Sprintf("Wrong response when reading struct bytes: %v", err))
	// }
	payload.FromBytes(readPayloadBuffer)
	return &message.Message{
		Type:    typeValue,
		Payload: payload,
	}, nil
}

func (s *SocketReceiver) ReadString() (string, error) {
	return s.reader.ReadString(s.Delim)
}

func NewSocketReceiver(connection io.Reader) *SocketReceiver {
	delim := byte('\n')
	reader := bufio.NewReader(connection)
	return &SocketReceiver{
		reader: reader,
		Delim:  delim,
	}
}

func (s *SocketReceiver) readType(connection io.Reader) (uint16, error) {
	readTypeBuffer, err := s.readNBytes(connection, message.TYPE_HEADER_SIZE_BYTES)
	fmt.Println("ReadTypeBuffer contents: ", readTypeBuffer)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("Could not read type bytes. Reading failed: %v", err))
	}
	return message.DecodeValue[uint16](readTypeBuffer)
}

func (s *SocketReceiver) readNBytes(connection io.Reader, n int) ([]byte, error) {
	scanner := bufio.NewScanner(connection)
	buffer := make([]byte, n)
	scanner.Buffer(buffer, len(buffer))
	for {
		if !scanner.Scan() {
			break
		}
	}
	return buffer, scanner.Err()
}
