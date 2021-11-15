package main

import (
	"server/internal/listener"
)

func main() {
	var server listener.IService
	server = &listener.Service{}
	server.RunServer()
}
