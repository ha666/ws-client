package main

import (
	"bytes"

	"github.com/ha666/logs"
	"github.com/ha666/ws-common"
	"github.com/maurodelazeri/gorilla-reconnect"
)

//服务器
const addr = "wss://ha666.com/echo"

//本机，需要绑定hosts
//const addr = "ws://websocket.com/echo"

func main() {
	conn := &recws.RecConn{}
	conn.Dial(addr, nil)
	for {
		receive(conn)
	}
}

func receive(conn *recws.RecConn) {
	dst, messageType, err := ws_common.ReadMessage(conn)
	if err != nil {
		logs.Error("read err:", err)
		return
	}
	if bytes.Compare(messageType, ws_common.MESSAGEPING) == 0 {
		Ping(conn, dst)
	} else if bytes.Compare(messageType, ws_common.MESSAGEPONG) == 0 {
		Pong(conn, dst)
	} else if bytes.Compare(messageType, ws_common.MESSAGEREAD) == 0 {
		Read(conn, dst)
	} else if bytes.Compare(messageType, ws_common.MESSAGEWRITE) == 0 {
		Write(conn, dst)
	} else if bytes.Compare(messageType, ws_common.MESSAGESUBSCRIPTION) == 0 {
		Subscription(conn, dst)
	} else if bytes.Compare(messageType, ws_common.MESSAGEPUBLISH) == 0 {
		Publish(conn, dst)
	} else {
		logs.Error("无效的消息类型")
	}
}
