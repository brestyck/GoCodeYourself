package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Адрес напиши чучело")
		os.Exit(14)
	}

	remoteAddr, _ := net.ResolveUDPAddr("udp", os.Args[1])
	conn, err := net.ListenPacket("udp", "")

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	// buffer := make([]byte, 1414)
	message := `
    ██████╗███╗   ██╗████████╗██████╗ ███████╗ █████╗ 
   ██╔════╝████╗  ██║╚══██╔══╝██╔══██╗██╔════╝██╔══██╗
   ██║     ██╔██╗ ██║   ██║   ██████╔╝█████╗  ███████║
   ██║     ██║╚██╗██║   ██║   ██╔══██╗██╔══╝  ██╔══██║
   ╚██████╗██║ ╚████║   ██║   ██║  ██║███████╗██║  ██║
    ╚═════╝╚═╝  ╚═══╝   ╚═╝   ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝											  
   `

	for {
		bytesRead, _ := conn.WriteTo([]byte(message), remoteAddr)
		if err != nil {
			log.Fatalf("O KURWA OSHIBKAAAAAAAAAAAA: %v", err)
		}
		log.Printf("Sent: %v bytes from %v", bytesRead, remoteAddr)
	}
}
