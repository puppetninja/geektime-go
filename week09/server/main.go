package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/puppetninja/geektime-go/week09/goim"
)

// Unhandled tcp stream bytes
func handleConn(conn net.Conn) {
	defer conn.Close()
	defer fmt.Println("close tcp connection")
	fmt.Println("new tcp connection: ", conn.RemoteAddr())

	result := bytes.NewBuffer(nil)
	var buf [1024]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				continue
			} else {
				fmt.Println("read err:", err)
				break
			}
		} else {
			fmt.Println("recv:", result.String())
		}
		result.Reset()
	}
}

// Handle tcp steam bytes with goim frame.
func handleGoimConn(conn net.Conn) {
	defer conn.Close()
	defer fmt.Println("close tcp connection")
	fmt.Println("new tcp connection: ", conn.RemoteAddr())
	result := bytes.NewBuffer(nil)
	// setup scanner
	scanner := bufio.NewScanner(result)
	scanner.Split(goim.SplitFunc)

	for scanner.Scan() {
		msg, _ := goim.Decode(bytes.NewReader(scanner.Bytes()))
		log.Println(msg)
	}
	// Meh
	if err := scanner.Err(); err != nil {
		log.Fatal("Invalid frame")
	}
}

func main() {
	l, err := net.Listen("tcp", ":31000")
	if err != nil {
		panic(err)
	}
	fmt.Print("Listening on port 31000")

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("conn err:", err)
		} else {
			go handleGoimConn(conn)
		}
	}
}
