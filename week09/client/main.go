package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/puppetninja/geektime-go/week09/goim"
)

func main() {
	conn, err := net.DialTimeout("tcp", "localhost:31000", time.Second*30)
	if err != nil {
		fmt.Printf("connect failed, err : %v\n", err.Error())
		return
	}

	msg := goim.Message{
		Body: []byte("[GOIM frame full message.]"),
	}
	data, err := goim.Encode(&msg)
	if err != nil {
		fmt.Printf("encode failed, err : %v\n", err.Error())
		return
	}
	log.Printf(string(data))

	for i := 0; i < 1000; i++ {
		_, err = conn.Write(data)
		if err != nil {
			fmt.Printf("write failed , err : %v\n", err)
			break
		}
	}
}
