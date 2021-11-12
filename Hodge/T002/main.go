package main

import (
	"bufio"
	"fmt"
	"net"
)

type EntireTick struct {
	Date      string  `json:"Date"`
	Time      string  `json:"Time"`
	TickType  int64   `json:"TickType"`
	AmountSum float64 `json:"AmountSum"`
	Close     float64 `json:"Close"`
	Volume    int64   `json:"Volume"`
	VolumeSum int64   `json:"VolumeSum"`
	Topic     string  `json:"Topic"`
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [256]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Printf("read from conn failed, err:%v\n", err)
			break
		}

		recv := string(buf[:n])
		fmt.Printf("%v\n", recv)

		_, err = conn.Write([]byte("ok"))
		if err != nil {
			fmt.Printf("write from conn failed, err:%v\n", err)
			break
		}
	}
}

func main() {
	listen, err := net.Listen("tcp", "172.20.10.105:6666")
	if err != nil {
		fmt.Printf("listen failed, err:%v\n", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed, err:%v\n", err)
			continue
		}
		go process(conn)
	}
}
