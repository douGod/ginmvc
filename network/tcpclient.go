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
var isChannelClose bool
func SetUpTcpClient(){
	tcponce.Do(func(){
		msgChannel = make(chan []byte,1000)
		closeChannel = make(chan []byte,1)
		isChannelClose = false
		tcpconn,err = net.Dial("tcp","127.0.0.1:9502")
		defer tcpconn.Close()
		if err != nil{
			log.Fatal(err)
		}
		//一直向服务端发送消息
		var data []byte
		for {
			select{
				case data = <- msgChannel:
					if _,err = tcpconn.Write(data);err != nil{
						goto ERR
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
		isChannelClose = true
		close(closeChannel)
	})
}
func SendMessage(message []byte){
	if isChannelClose{
		log.Fatal("消息通道已经关闭")
	}
	msgChannel <- message

}
