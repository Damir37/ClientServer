package main

import (
	"ClientChat/encrypt"
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func ConnectServer() (string, net.Conn) {
	connect, err := net.Dial("tcp", "127.0.0.1:3000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Client connect!")
	buff := make([]byte, 1024)
	data, err := connect.Read(buff)
	if err != nil {
		log.Fatal(err)
	}
	succ := string(buff[0:data])
	return succ, connect
}

func sendMessage(msg string, connect net.Conn) {
	write, err := connect.Write([]byte(msg))
	if err != nil || write == 0 {
		log.Fatal(err)
	}
}

func main() {
	date, connect := ConnectServer()
	fmt.Println(date)
	for {
		buf := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		msg, err := buf.ReadBytes('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(encrypt.EcnryptMessage(string(msg)))
		sendMessage(encrypt.EcnryptMessage(string(msg)), connect)
	}

}
