package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Set a read and write deadline for the connection
	// conn.SetDeadline(time.Now().Add(10 * time.Second))

	message, err := receiveMessage(conn)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Received message:", message)

	response := "Hello from the server!"
	err = sendMessage(conn, response)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func receiveMessage(conn net.Conn) (string, error) {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		return "", err
	}
	return string(buf[:n]), nil
}

func sendMessage(conn net.Conn, message string) error {
	_, err := conn.Write([]byte(message))
	return err
}
