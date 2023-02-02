package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.ListenPacket("udp", ":1414")

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	buffer := make([]byte, 65536)

	for {
		bytesRead, remoteAddr, err := conn.ReadFrom(buffer)
		if err != nil {
			log.Printf("O KURWA OSHIBKAAAAAAAAAAAA: %v", err)
		}
		log.Printf("Read: %v bytes from %v", bytesRead, remoteAddr)
		fmt.Println(string(buffer))
	}
}
