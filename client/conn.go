package client

import (
	"bufio"
	"net"
	"time"
)

type Connection struct {
	netcn net.Conn
	rd    *bufio.Reader
	buf   []byte
}

func NewConnection() (*Connection, error) {
	conn, err := net.DialTimeout("tcp", "", 5*time.Second)
	if err != nil {
		return nil, err
	}

	cn := &Connection{
		netcn: conn,
		buf:   make([]byte, 0),
	}

	cn.rd = bufio.NewReader(cn)
	return cn, nil
}

func (c Connection) Read(b []byte) (n int, err error) {
	return c.netcn.Read(b)
}

func (c Connection) Write(b []byte) (n int, err error) {
	return c.netcn.Write(b)
}

func (c *Connection) Close() {
	c.netcn.Close()
}
