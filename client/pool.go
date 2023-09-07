package client

import (
	"sync"
)

type ConnPool struct {
	pool chan Connection
	mu   sync.Mutex
}

func NewConnectionPool() *ConnPool {
	pool := make(chan Connection, numClients)
	return &ConnPool{
		pool: pool,
	}
}

func (p *ConnPool) Get() (Connection, error) {
	select {
	case conn := <-p.pool:
		return conn, nil
	default:
		conn, err := NewConnection()
		if err != nil {
			return Connection{}, err
		}
		return *conn, nil
	}
}

func (p *ConnPool) Put(conn Connection) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.pool <- conn
}
