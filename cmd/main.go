package main

import (
	_ "server/internal/battleground"
	"server/internal/listener"
)

func main() {
	var server listener.IService = &listener.Service{}
	server.RunServer()
}
