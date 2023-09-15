package server

import (
	"log"
	"net"

	"github.com/ddonskaya/feather/collections"
	"github.com/ddonskaya/feather/utils"
)

type Server struct {
	buffers  *utils.BufferPool
	hm       collections.HashMapCollection
	ss       collections.SortedSetCollection
	connPool ConnPool
	options  *Options
	logger   *log.Logger
}

func FeatherServer(opt *Options) *Server {
	s := &Server{
		buffers:  utils.NewBufferPool(),
		hm:       *collections.NewHashMapCollection(),
		ss:       *collections.NewSortedSetCollection(),
		connPool: *NewConnectionPool(),
		options:  opt,
		logger:   opt.setLogger(),
	}
	s.HandleTCP()
	return s
}

func (s *Server) HandleTCP() error {
	s.logger.Println("Start Feather server")

	connTCP, err := net.ListenTCP("tcp", s.options.GetTCPAddress())
	if err != nil {
		s.logger.Printf("server: can not listen tcp connection %v", err)
		return err
	}

	for {
		conn, err := connTCP.AcceptTCP()
		s.connPool.Put(conn)
		if err != nil {
			log.Println(err)
		}

		if err := conn.SetNoDelay(false); err != nil {
			log.Println(err)
			conn.Close()
			continue
		}

		go newClient(conn, s).ProcessCmd()
	}
}