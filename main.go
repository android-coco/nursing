package main

import (
	_ "nursing/routers"
	_"nursing/utils"
	"fit"
)
func main() {
	defer func() {
		fit.Stop()
	}()
	fit.Start() //开始监听
}
