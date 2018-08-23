package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	//returns a listener, in our case a tcp listener
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T\n", listener)
	//Handle connection
	for {
		//Accepts and creates a connection
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn) //Handles one connection at a time
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
