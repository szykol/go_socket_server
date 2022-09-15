package main

import (
	"fmt"
	"net"
	"szykol/sockets/pkg/receiver"
	// "szykol/sockets/pkg/sender"
)

func main() {
	acceptor, err := net.Listen("tcp", ":1338")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer acceptor.Close()
	fmt.Println("Acceptor initialized")

	for {
		connection, err := acceptor.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("New connection acquired")

		go handle(connection)
	}
}

func handle(con net.Conn) {
	fmt.Println("[server] Handling new connection")
	receiver := receiver.NewSocketReceiver(con)
	// sender := sender.NewSocketSender(con)
	for {
		data, err := receiver.ReadMessage(con)
		if err != nil {
			fmt.Println("[server] err: ", err)
			break
		}
		fmt.Println("[server] Received: ", data.Payload)
		fmt.Printf("[server] Received following data: %v", data)
		con.Close()
		break
	}
	con.Close()
}
