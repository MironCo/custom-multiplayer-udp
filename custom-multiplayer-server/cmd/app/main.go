package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	address, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		log.Fatal("Error resolving UDP Address")
	}

	connection, err := net.ListenUDP("udp", address)
	if err != nil {
		log.Fatal("Error starting UDP server:", err)
	}
	defer connection.Close()

	fmt.Println("UDP Server listing on Port http://127.0.0.1:8080")

	buffer := make([]byte, 1024)
	for {
		n, clientAddress, err := connection.ReadFromUDP(buffer)
		if err != nil {
			log.Println("Error reading UDP message: ", err)
			continue
		}

		message := string(buffer[:n])
		fmt.Println("Received from %s: %s\n", clientAddress, message)

		response := fmt.Sprintf("Echo: %s", message)
		_, err = connection.WriteToUDP([]byte(response), clientAddress)
		if err != nil {
			log.Println("Error sending response: ", err)
		}
	}
}
