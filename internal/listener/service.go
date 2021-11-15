package listener

import (
	"fmt"
	"net"
)

type IService interface {
	RunServer()
}

type Service struct {
	listener listener
}

func handleRequest(connection net.Conn) {
	// Здесь пишем обработчик соединения
	fmt.Println(connection)
}

func (s *Service) RunServer() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
		return
	}
	fmt.Println("Server is running localhost:8080")
	s.listener.server = ln
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			return
		}
		// Что-то делаем
		go handleRequest(conn)
	}
}

func (s *Service) StopServer() {
	err := s.listener.server.Close()
	if err != nil {
		// handle error
		return
	}
	fmt.Println("TCP server has been successfuly stopped")
}
