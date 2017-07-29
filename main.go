package main

import (
	"runtime"

	"github.com/cocobao/shitcake/server"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	server.Setup()
}
