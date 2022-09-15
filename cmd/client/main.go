package main

import (
	// "bufio"
	"fmt"
	"net"
	"szykol/sockets/pkg/message"
	"szykol/sockets/pkg/receiver"
	"szykol/sockets/pkg/sender"
)

func main() {
	connection, err := net.Dial("tcp", "0.0.0.0:1338")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer connection.Close()
	// msg := "hello, world\n"
	// fmt.Printf("[client] Sending following data: %v", msg)
	// newWriter := bufio.NewWriter(connection)
	// newWriter.Flush()
	sender := sender.NewSocketSender(connection)
	payload := message.VectorPayload{
		X: 2,
		Y: 3,
	}
	msg := message.Message{
		Type:    message.MSG_T_EXAMPLE,
		Payload: &payload,
	}
	bytes, err := msg.ToBytes()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("[client] Sending following bytes: ", bytes)
	sender.SendBytes(bytes)

	receiver := receiver.NewSocketReceiver(connection)
	data, err := receiver.ReadString()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("[client] Received following data: %v", data)
}
