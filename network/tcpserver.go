package network

import (
	"fmt"
	"log"
	"net"
	"sync"
)
var mapTcpConn map[string] net.Conn
var tcpservonce sync.Once
func SetUpTcpServer(){
	tcpservonce.Do(func(){
		mapTcpConn = make(map[string] net.Conn)
		list,err := net.Listen("tcp","0.0.0.0:9502")
		fmt.Println("success to setup tcp serv")
		defer list.Close()
		if err != nil{
			log.Fatal(err)
		}
		for {
			conn,err := list.Accept()
			if err != nil{
				continue
			}
			go delWithConn(conn)
		}
	})
}

func delWithConn(conn net.Conn){
	mapTcpConn[conn.RemoteAddr().String()] = conn
	data := make([]byte,2048)
	var err error
	var n int
	for{
		n,err = conn.Read(data)
		if err != nil{
			fmt.Println("failed to read from client")
			return
		}
		fmt.Printf("recieve message:%s \r\n",string(data[:n]))
		_,err = conn.Write(data[:n])
		if err != nil{
			fmt.Println("failed to respond to  client")
			return
		}
	}
}