package main

import (
	"fmt"
	"server/internal/listener"
	_ "server/internal/battleground"
)

func main() {
	gay := new(int)
	fmt.Println(*gay)
	var server listener.IService = &listener.Service{}
	go server.RunServer()
	for {

	}
}
