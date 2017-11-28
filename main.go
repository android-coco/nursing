package main

import (
	_ "nursing/routers"
	_"nursing/background"
	"fit"
)
func main() {
	defer func() {
		fit.Stop()
	}()
	fit.Start() //开始监听
}
