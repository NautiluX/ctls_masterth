package main

import (
	"fmt"
	"net"
	"os"
)

const (
	CONN_HOST = "0.0.0.0"
	CONN_PORT = "1234"
	CONN_TYPE = "tcp"
)

func main() {
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	fmt.Println("New Client connected!")
	bytes := readClientHello(conn)
	if bytes == 132 {
		sendServerHelloTls(conn)
	}
	if bytes == 50 {
		sendServerHelloCtls(conn)
	}
	readClientHello(conn)
	// Close the connection when you're done with it.
	conn.Close()
	fmt.Println("Connection closed!")
}

func readClientHello(conn net.Conn) int {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Printf("%d bytes received: %v\n", reqLen, string(buf))
	return reqLen

}

func sendServerHelloTls(conn net.Conn) {
	// Send a response back to person contacting us.
	shello := ""
	for i := 0; i < 90+478; i++ {
		shello += "h"
	}
	conn.Write([]byte(shello))
}
func sendServerHelloCtls(conn net.Conn) {
	// Send a response back to person contacting us.
	shello := ""
	for i := 0; i < 152; i++ {
		shello += "h"
	}
	conn.Write([]byte(shello))
}
