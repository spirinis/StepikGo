package network

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

// Подключитесь к адресу 127.0.0.1:8081 по протоколу TCP, считайте от сервера 3 сообщения, и выведите их в верхнем регистре.
func ConnectionTCP() {
	go Server()
	time.Sleep(time.Millisecond * 10)
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()
	// Таймаут на первое чтение
	conn.SetReadDeadline(time.Now().Add(time.Second))
	for range 3 {
		message := make([]byte, 1024)
		nRead, err := conn.Read(message)
		if err != nil {
			log.Println(err)
			break
		}
		// Таймаут на последующие чтения
		conn.SetReadDeadline(time.Now().Add(time.Millisecond * 500))
		fmt.Println(strings.ToUpper(string(message[:nRead])))
		//time.Sleep(time.Millisecond * 10)
	}
}

func Server() {
	listener, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Println(err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	conn.Write([]byte("message"))

	time.Sleep(time.Millisecond * 1500)

	conn.Write([]byte("MesSaGe"))

	time.Sleep(time.Millisecond * 1500)

	conn.Write([]byte("MESSAGE"))

}
