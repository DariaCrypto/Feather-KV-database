package server

import (
	"net"
	"sync"
)

const (
	serverAddress = "127.0.0.1:6870"
	numClients    = 10
)

type ConnPool struct {
	pool chan net.Conn
	mu   sync.Mutex
}

func NewConnectionPool() *ConnPool {
	pool := make(chan net.Conn, numClients)
	return &ConnPool{
		pool: pool,
	}
}

func (p *ConnPool) Get() (net.Conn, error) {
	select {
	case conn := <-p.pool:
		return conn, nil
	default:
		conn, err := net.Dial("tcp", serverAddress)
		if err != nil {
			return nil, err
		}
		return conn, nil
	}
}

func (p *ConnPool) Put(conn net.Conn) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.pool <- conn
}
