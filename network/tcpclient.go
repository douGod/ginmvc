package network

import (
	"log"
	"net"
	"sync"
)
var tcpconn net.Conn
var tcponce sync.Once
var err error
var closeOnce sync.Once
var msgChannel chan []byte
var closeChannel chan []byte
func SetUpTcpClient(){
	tcponce.Do(func(){
		msgChannel = make(chan []byte,1000)
		closeChannel = make(chan []byte,1)
		tcpconn,err = net.Dial("tcp","127.0.0.1:9502")
		defer tcpconn.Close()
		if err != nil{
			log.Fatal(err)
		}
		//模拟一直向服务端发送消息
		var data []byte
		for {
			select{
				case data = <- msgChannel:
					if _,err = tcpconn.Write(data);err != nil{
						continue
					}
				case <-closeChannel:
					goto ERR
			}
		}
		ERR:
			closeChan()
	})

}
func closeChan(){
	closeOnce.Do(func(){
		close(closeChannel)
	})
}
func SendMessage(message []byte){
	msgChannel <- message
}
