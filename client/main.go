package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/sencoder/SocketBenchmark/client/util"
)

const (
	SERVER_ADDRESS = "192.168.123.163:9090"
)

type Client struct {
	conn *net.TCPConn
}

func (cli *Client) Start() (err error) {

	tcpAddr, err := net.ResolveTCPAddr("tcp4", SERVER_ADDRESS)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		return
	}

	cli.conn, err = net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		return
	}
	defer cli.conn.Close()
	log.Println("connect success")

	work(cli.conn)

	log.Println("client finish")
	return nil
}

// 收取 socket server 信息
func work(conn *net.TCPConn) error {
	buffer := make([]byte, 1024)

	for {
		conn.Write([]byte("Pride and Prejudice. It is a truth universally acknowledged, that a single man in possession of a good fortune must be in wanna of a wife."))
		startTime := time.Now().UnixNano() / 1000000
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println(conn.RemoteAddr().String(), "connection error:", err)
			break
		} else {
			msg := string(buffer[:n])
			log.Println("Msg:", msg)
		}
		endTime := time.Now().UnixNano() / 1000000
		s := util.Sample{Time: time.Now().Unix(), Latency: endTime - startTime}
		util.WriteData(s)
	}
	return nil
}

func addClient() {

	go new(Client).Start()
}

func main() {
	cliNum := 60
	for i := 0; i < cliNum; i++ {
		time.Sleep(time.Second)
		addClient()
		log.Println("Add client", i)
	}
}
