package server

import (
	"fmt"
	"log"
	"net"
	"sync"
	"github.com/ddonskaya/feather/utils"
	"github.com/ddonskaya/feather/collections"

)

const (
	listenAddress = "127.0.0.1:8080"
	poolSize      = 5
)

type Server struct {
	buffers *utils.BufferPool
	hm collections.HashMapCollection
	ss collections.SortedSetCollection
}

func FeatherServer() *Server {
	return &Server{
		buffers: utils.NewBuffer(),
		hm: *collections.NewHashMapCollection(),
		ss: collections.SortedSetCollection{},
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

		go newClient(conn).HandleClientCmd()
	}
}