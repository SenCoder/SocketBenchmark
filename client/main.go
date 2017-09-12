package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/sencoder/SocketBenchmark/client/util"
)

const (
	SERVER_ADDRESS = "13.56.82.138:9000"
)

var c util.Collector

type Client struct {
	conn  *net.TCPConn
	count int64
}

func init() {
	c.OpenFile("sample.json", 0644)
}

func (cli *Client) Start() (err error) {

	tcpAddr, err := net.ResolveTCPAddr("tcp4", SERVER_ADDRESS)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		return
	}

	cli.conn, err = net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		return
	}
	defer cli.conn.Close()
	log.Println("connect success")

	cli.work(cli.conn)

	log.Println("client finish")
	return nil
}

// 收取 socket server 信息
func (cli *Client) work(conn *net.TCPConn) error {
	buffer := make([]byte, 1024)

	msg := util.Rs(512)

	for {
		startTime := time.Now().UnixNano()
		conn.Write(msg)

		_, err := conn.Read(buffer)
		if err != nil {
			log.Println(conn.RemoteAddr().String(), "connection error:", err)
			break
		} else {
			_, err = conn.Write([]byte(msg))
			if err != nil {
				log.Println(conn.RemoteAddr().String(), "connection error:", err)
				break
			}
		}
		cli.count++
		endTime := time.Now().UnixNano()
		d := util.DataSample{Time: time.Now().Unix(), Latency: (endTime - startTime) / 1000000, Count: cli.count}
		c.Sample(d)
		time.Sleep(time.Second)
	}
	return nil
}

func addClient() {

	go new(Client).Start()
}

func main() {

	flag := make(chan bool)

	if len(os.Args) < 2 {
		log.Println("Please give an int parameter")
		return
	}

	arg := os.Args[1]

	cliNum, err := strconv.Atoi(arg)
	if err != nil {
		log.Println(err)
		return
	}

	for i := 0; i < cliNum; i++ {
		time.Sleep(time.Millisecond * 50)
		addClient()
		log.Println("Add client", i)
	}
	flag <- true
}
