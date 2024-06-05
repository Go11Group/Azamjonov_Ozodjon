package main

import (
	"fmt"
	"net"
)

func main() {
	serverAddr, err := net.ResolveUDPAddr("udp", "localhost:9090")
	if err != nil {
		fmt.Println("Server manzilini aniqlashda xato:", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Println("Serverga ulanishda xato:", err)
		return
	}
	defer conn.Close()

	go receiveMessages(conn)

	message := "Salom, server!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Xabarni yuborishda xato:", err)
		return
	}

	// Foydalanuvchi kiritishlarini o'qish va yuborish
	for {
		fmt.Print("Yozing: ")
		var input string
		fmt.Scanln(&input)
		_, err = conn.Write([]byte(input))
		if err != nil {
			fmt.Println("Xabarni yuborishda xato:", err)
			return
		}
	}
}

func receiveMessages(conn *net.UDPConn) {
	buffer := make([]byte, 1024)
	for {
		n, _, err := conn.ReadFrom(buffer)
		if err != nil {
			fmt.Println("Serverdan javobni o'qishda xato:", err)
			return
		}
		fmt.Printf("Serverdan javob: %s\n", string(buffer[:n]))
	}
}
