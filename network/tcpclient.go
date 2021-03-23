package network

import (
	"log"
	"net"
	"time"
)

func SetUpTcpClient(){
	var conn net.Conn
	var err error
	conn,err = net.Dial("tcp","127.0.0.1:9502")
	if err != nil{
		log.Fatal(err)
	}
	//模拟一直向服务端发送消息
	for {
		time.Sleep(time.Second * 1)
		_,err = conn.Write([]byte("我爱你"))
		if err != nil{
			log.Fatal(err)
		}
	}

}
