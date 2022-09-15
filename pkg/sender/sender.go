package sender

import (
	"io"
)

type Sender interface {
	SendBytes([]byte) error
	SendString(string) error
}

type SocketSender struct {
	connection io.Writer
}

func (s *SocketSender) SendBytes(msg []byte) error {
	_, err := s.connection.Write(msg)
	return err
}

func (s *SocketSender) SendString(msg string) error {
	return s.SendBytes([]byte(msg))
}

func NewSocketSender(connection io.Writer) *SocketSender {
	return &SocketSender{
		connection: connection,
	}
}
