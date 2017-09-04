package main

import (
	"log"
	"net"
	"runtime"
	"sync"
	"time"

	"github.com/sencoder/SocketBenchmark/server/util"
)

var id = 1000

var plock chan bool

var clientMgr ClientMgr

func init() {
	plock = make(chan bool, 2)
	clientMgr.clients = make(map[int]*Client)
	clientMgr.length = 0
}

func lock() {
	plock <- true
}

func unlock() {
	<-plock
}

func StartSocketServer(host string) {
	netListen, err := net.Listen("tcp", host)
	if err != nil {
		log.Println(err)
		return
	}
	defer netListen.Close()

	for {
		log.Println("Waiting for clients")
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}
		log.Println(conn.RemoteAddr().String(), " tcp connect success")

		cli := Client{Id: getId()}
		clientMgr.AddClient(cli)

		go cli.onConnect(conn)
	}
}

func getId() int {
	lock()
	defer unlock()
	id++
	return id
}

type Client struct {
	Id int
}

type ClientMgr struct {
	lock    sync.Mutex
	clients map[int]*Client
	length  int
}

func (mgr *ClientMgr) AddClient(client Client) {
	mgr.lock.Lock()
	defer mgr.lock.Unlock()

	mgr.clients[client.Id] = &client
	mgr.length++
	log.Println("clients:", mgr.length)
}

func (mgr *ClientMgr) DeleteClient(client Client) {
	mgr.lock.Lock()
	defer mgr.lock.Unlock()

	delete(mgr.clients, client.Id)
	mgr.length--
	log.Println("clients:", mgr.length)
}

//handle the connection
func (cli *Client) onConnect(conn net.Conn) {

	defer conn.Close()
	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println(conn.RemoteAddr().String(), " connection error:", err)
			cli.onDisConnect()
			return
		}
		time.Sleep(time.Millisecond * 5)
		msg := "echo: " + string(buffer[:n])

		n, err = conn.Write([]byte(msg))
		if err != nil {
			log.Println("Seem to send msg fail:", err)
		}
	}
}

func (cli *Client) onDisConnect() {
	clientMgr.DeleteClient(*cli)
}

func main() {
	runtime.GOMAXPROCS(6)

	util.CollectData()
	StartSocketServer("0.0.0.0:9090")
}
