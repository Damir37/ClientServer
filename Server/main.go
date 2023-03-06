package main

import (
	"ServerChat/encrypt"
	"fmt"
	"log"
	"net"
)

func ServerStart() {
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server started")
	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			listener.Close()
		}
		go handleConnection(connection)
	}
}

func handleConnection(connect net.Conn) {
	succ := "Succ connect!"
	connect.Write([]byte(succ))
	buff := make([]byte, 1024)
	datein, err := connect.Read(buff)
	if err != nil || datein == 0 {
		log.Fatal(err)
	}
	message := string(buff[0:datein])
	fmt.Println("ENCODED:", message)
	fmt.Println("DECODED:", encrypt.DecryptMessage(message))
}

func main() {
	ServerStart()
}
