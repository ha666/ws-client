package main

import "github.com/ha666/logs"

func init() {
	initLog()
}

func initLog() {
	_ = logs.SetLogger(logs.AdapterConsole, `{"level":7}`)
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)
}
