package main

import (
	"fmt"
	"net"
	"sync"
)

var clients = make(map[string]net.Addr)
var mutex = &sync.Mutex{}

func main() {
	conn, err := net.ListenPacket("udp", ":9090")
	if err != nil {
		fmt.Println("Tinglashni sozlashda xato:", err)
		return
	}
	defer conn.Close()
	fmt.Println("Server 9090-portda tinglamoqda...")

	buffer := make([]byte, 1024)

	for {
		n, addr, err := conn.ReadFrom(buffer)
		if err != nil {
			fmt.Println("Xabarni o'qishda xato:", err)
			continue
		}

		message := string(buffer[:n])
		fmt.Printf("Qabul qilingan xabar: %s\n", message)
		addClient(addr)
		broadcastMessage(conn, addr, message)
	}
}

func addClient(addr net.Addr) {
	mutex.Lock()
	defer mutex.Unlock()
	clients[addr.String()] = addr
}

func broadcastMessage(conn net.PacketConn, senderAddr net.Addr, message string) {
	mutex.Lock()
	defer mutex.Unlock()
	for _, clientAddr := range clients {
		if clientAddr.String() != senderAddr.String() {
			response := fmt.Sprintf("%s => %s", senderAddr.String(), message)
			_, err := conn.WriteTo([]byte(response), clientAddr)
			if err != nil {
				fmt.Println("Xabarni yozishda xato:", err)
			}
		}
	}
}
