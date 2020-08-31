package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/ha666/logs"
	"github.com/ha666/ws-common/protocol"
	"github.com/maurodelazeri/gorilla-reconnect"
)

func Ping(c *recws.RecConn, dst proto.Message) {
	val, ok := dst.(*protocol.Ping)
	if !ok {
		logs.Error("解析ping消息出错")
		return
	}
	logs.Info("\tmessageType:%s\tmessage: %s", "ping", val.PingVal)
}

func Pong(c *recws.RecConn, dst proto.Message) {
	val, ok := dst.(*protocol.Pong)
	if !ok {
		logs.Error("解析pong消息出错")
		return
	}
	logs.Info("\tmessageType:%s\tmessage: %s", "pong", val.PongVal)
}

func Read(c *recws.RecConn, dst proto.Message) {
	val, ok := dst.(*protocol.Read)
	if !ok {
		logs.Error("解析read消息出错")
		return
	}
	logs.Info("\tmessageType:%s\tmessage: %s", "read", val.ReadVal)
}

func Write(c *recws.RecConn, dst proto.Message) {
	val, ok := dst.(*protocol.Write)
	if !ok {
		logs.Error("解析write消息出错")
		return
	}
	logs.Info("\tmessageType:%s\tmessage: %s", "write", val.WriteVal)
}

func Subscription(c *recws.RecConn, dst proto.Message) {
	val, ok := dst.(*protocol.Subscription)
	if !ok {
		logs.Error("解析subscription消息出错")
		return
	}
	logs.Info("\tmessageType:%s\tmessage: %s", "subscription", val.SubscriptionVal)
}

func Publish(c *recws.RecConn, dst proto.Message) {
	val, ok := dst.(*protocol.Publish)
	if !ok {
		logs.Error("解析publish消息出错")
		return
	}
	logs.Info("\tmessageType:%s\tmessage: %s", "publish", val.PublishVal)
}
