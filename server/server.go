package server

import (
	"fmt"
	"log"
	"net"
	"sync"
)

const (
	listenAddress = "127.0.0.1:8080"
	poolSize      = 5
)

type Server struct {
	buffers *utils.BufferPool

}

func FeatherServer() *Server {
	return &Server{
		pool: pool,
	}
}

func (s *Server) HandleTCP() error {
	tcpLS, err :=  net.ListenTCP("tcp", listenAddress)
	if err != nil {
		return err
	}

	for {
		conn, err := tcpLS.AcceptTCP()
		if err != nil {
			log.Println(err)
		}

		if err := conn.SetNoDelay(false); err != nil {
			log.Println(err)
			conn.Close()
			continue
		}

		go newClient(conn, s).handleCommand()
	}
}