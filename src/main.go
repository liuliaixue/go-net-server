package main

import (
	"fmt"
	"net"
	// "time"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {

	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			fmt.Println(err)
			continue
		}
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	print( `start`)
	defer conn.Close()
	// err := conn.SetDeadline(time.Now().Add(time.Minute * 3))
	// if err != nil {
		
	// }
	var buf [1024]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			fmt.Println(n, err)
			return
		}
		n, err = conn.Write(buf[0:n])
		if err != nil {
			fmt.Println(n, err)
			return
		}
		fmt.Println(n ,buf)
		print(`end`)
	}

}
