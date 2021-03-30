package network

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)
var(
	Upgrade = websocket.Upgrader{
		//允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

//连接信息结构体
type conn struct {
	Connection *websocket.Conn
	BindID string
	IsClose bool
	Muetex sync.Mutex
	InChannel chan []byte
	OutChannel chan []byte
	CloseChan chan []byte
}
var mapConnect sync.Map
func init(){

}
//关闭连接
func (con *conn)CloseWs(){
	mapConnect.Delete(con.BindID)
	con.Connection.Close()
	con.Muetex.Lock()
	if !con.IsClose{
		con.IsClose = true
		close(con.CloseChan)
	}
	con.Muetex.Unlock()
}
//循环读取接收到的信息
func (con *conn)ReadLoop(){
	var data []byte
	var err error
	for{
		if _,data,err = con.Connection.ReadMessage();err != nil{
			fmt.Println(err)
			goto ERR
		}
		con.InChannel <- data
	}
	ERR:
		con.CloseWs()
}
//读取信息
func (con *conn)ReadMessage()(data []byte,err error ){
	//select多路IO复用
	select{
		case data = <- con.InChannel:
		case <- con.CloseChan:
			err = errors.New("failed to read message")
	}
	return
}
//循环写入信息
func (con *conn)WriteLoop(){
	//select多路IO复用
	var data []byte
	var err error
	for{
		select{
			case data = <- con.OutChannel:
				mapConnect.Range(func(k interface{},v interface{})bool{
					v1 := v.(*conn)
					if v1.BindID == con.BindID{
						if strings.Compare(string(data),"heartbeat") == 0{
							if err = v1.Connection.WriteMessage(websocket.TextMessage,data);err != nil{
								fmt.Println(err)
								con.CloseWs()
								return false
							}
						}
					}else{//别人发送的消息除了心跳都发送
						if strings.Compare(string(data),"heartbeat") != 0{
							if err = v1.Connection.WriteMessage(websocket.TextMessage,data);err != nil{
								fmt.Println(err)
								con.CloseWs()
								return false
							}
						}
					}
					return true
				})

			case <- con.CloseChan:
				goto ERR
		}
	}
ERR:
	con.CloseWs()
}
//发送信息
func (con *conn)WriteMessage(data []byte)(err error){
	select{
		case con.OutChannel <- data:
		case <-con.CloseChan:
			err = errors.New("failed to write message")
	}
	return
}
func WeChat(c *gin.Context){
	var err error
	var data []byte
	var Conn = &conn{}
	Conn.CloseChan = make(chan []byte,1)//缓冲区1个字节
	Conn.InChannel = make(chan []byte,1000)//缓冲区1000个字节
	Conn.OutChannel = make(chan []byte,1000)//缓冲区1000个字节
	Conn.IsClose = false
	Conn.BindID = c.Request.RemoteAddr
	if Conn.Connection,err = Upgrade.Upgrade(c.Writer,c.Request,nil);err != nil{
		Conn.CloseWs()
		log.Fatal(err)
		return
	}
	//存储链接
	mapConnect.Store(c.Request.RemoteAddr,Conn)
	//挂起循环读协程
	go Conn.ReadLoop()
	//挂起循环写协程
	go Conn.WriteLoop()

	//心跳检测
	go func(){
		for{
			time.Sleep(time.Second * 5)
			if err = Conn.WriteMessage([]byte("heartbeat"));err != nil{
				return
			}
		}
	}()
	for{
		//读数据
		if data,err = Conn.ReadMessage();err != nil{
			fmt.Println(err)
			return
		}
		//写数据
		if err = Conn.WriteMessage(data);err != nil{
			fmt.Println(err)
			return
		}
	}
}


