package main

import (
	// "encoding/binary"
	"fmt"
	"net"
	// "time"
)

func main() {
	fmt.Println("8080")
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
	fmt.Println(`start`)
	defer conn.Close()
	// err := conn.SetDeadline(time.Now().Add(time.Minute * 3))
	// if err != nil {

	// }

	var buf [512]byte
	var bufs [][]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			fmt.Println("connection error", n, err)
			return
		}
		bufs = append(bufs, buf[0:n])
		onBuffer(bufs)

		//n, err = conn.Write(buf[0:n])
		//if err != nil {
		//	fmt.Println(n, err)
		//	return
		//}

	}
	fmt.Println(`end`)

}
func getByteLength(buf []byte) int{
	var length int
	for index, value := range buf {
		p := int(1 << (uint32(index) * 8))
		length +=   (int(value) * p)
	}
	return length
}

func onBuffer(bufs [][]byte) {
	if len(bufs[0]) == 4 {
		fmt.Println("laji")
		
		lengthByte := bufs[0][0:4]
		fmt.Println(lengthByte)
		length := getByteLength(lengthByte)
		fmt.Println("length:", length)
	}

	if len(bufs[0]) > 4 {

		lengthByte := bufs[0][0:4]
		fmt.Println(lengthByte)
		length := getByteLength(lengthByte)
		fmt.Println("length:", length)
		if len(bufs[0]) > int(length+4) {
			data := bufs[0][:length+4]
			onData(data)
		}
	} else {
		if len(bufs) > 1 {
			firstBuf := bufs[0]
			bufs = bufs[1:]
			fmt.Println(firstBuf)
			fmt.Println(bufs)
			bufs[0] = append(firstBuf, bufs[0]...)
			onBuffer(bufs)
		}
	}
}

func onData(message []byte) {
	fmt.Println(message)
	fmt.Println(string(message))
}
