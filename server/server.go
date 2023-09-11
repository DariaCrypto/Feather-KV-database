package server

import (
	"log"
	"net"
	"github.com/ddonskaya/feather/utils"
	"github.com/ddonskaya/feather/collections"
)

type Server struct {
	buffers *utils.BufferPool
	hm collections.HashMapCollection
	ss collections.SortedSetCollection
	options *Options
	logger *log.Logger
}

func FeatherServer(opt *Options) *Server {
	s := &Server{
		buffers: utils.NewBuffer(),
		hm: *collections.NewHashMapCollection(),
		ss: *collections.NewSortedSetCollection(),
		options: opt,
		logger: opt.setLogger(),
	}
	s.HandleTCP()
	return s	
}


func (s *Server) HandleTCP() error {
	s.logger.Println("Start")
	tcpLS, err :=  net.ListenTCP("tcp", s.options.GetTCPAddress())
	if err != nil {
		s.logger.Printf("server: %v", err)
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

		go newClient(conn).ProcessCmd()
	}
}