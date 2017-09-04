package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

const (
	SERVER_ADDRESS = "192.168.123.163:9090"
)

var conn *net.TCPConn
var isStart bool

func init() {
	isStart = false
}

func StartService() (err error) {

	tcpAddr, err := net.ResolveTCPAddr("tcp4", SERVER_ADDRESS)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		return
	}

	conn, err = net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		return
	}
	defer conn.Close()
	log.Println("connect success")
	isStart = true

	receiveMsg()
	return nil
}

func StopService() {
	isStart = false
}

// 收取 socket server 信息
func receiveMsg() error {
	buffer := make([]byte, 1024)
	for isStart {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println(conn.RemoteAddr().String(), "connection error:", err)
			isStart = false
		} else {
			msg := string(buffer)
			log.Println("Msg:", msg, n)
		}
	}
	return nil
}

func main() {
	StartService()
}
