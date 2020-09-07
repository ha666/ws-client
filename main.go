package main

import (
	"bytes"
	"fmt"
	"time"

	"github.com/ha666/golibs"
	"github.com/ha666/logs"
	"github.com/ha666/ws-common"
	"github.com/ha666/ws-common/protocol"
	"github.com/maurodelazeri/gorilla-reconnect"
)

//服务器
//const addr = "ws://192.168.1.92:8888/process"

//本机，需要绑定hosts
const addr = "ws://websocket.com/process"

var (
	conn *recws.RecConn
)

func main() {
	for i := 0; i < 10; i++ {
		conn := &recws.RecConn{
			RecIntvlMin:      5 * time.Second,
			RecIntvlMax:      30 * time.Second,
			HandshakeTimeout: 10 * time.Millisecond,
		}
		conn.Dial(addr, nil)
		go read(conn)
		go ping(conn)
	}
	select {}
}

func read(conn *recws.RecConn) {
	for {
		dst, messageType, err := ws_common.ReadMessageWithAutoConnect(conn)
		if err != nil {
			logs.Error("read err:", err)
			time.Sleep(3 * time.Second)
			continue
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
}

func ping(conn *recws.RecConn) {
	for {
		time.Sleep(10 * time.Second)
		if err := ws_common.WriteMessageWithAutoConnect(conn, ws_common.MESSAGEPING, &protocol.Ping{
			PingVal: fmt.Sprintf("当前时间:%s", golibs.StandardTime()),
		}); err != nil {
			logs.Error("发送心跳失败:%s", err.Error())
		}
	}
}
