package server

import (
	"fmt"
	"net"

	"github.com/ddonskaya/feather/protocol"
	"github.com/ddonskaya/feather/utils"
	"google.golang.org/protobuf/proto"
)

type ConnHandler struct {
	conn      net.Conn
	server    *Server
	msgHeader []byte
}

func newClient(conn net.Conn, s *Server) *ConnHandler {
	return &ConnHandler{
		conn:      conn,
		msgHeader: make([]byte, utils.MSG_SIZE),
		server:    s,
	}
}

func (c *ConnHandler) ProcessCmd() {
	for {
		res, err := c.HandleClientCmd()
		if err != nil {
			c.server.logger.Printf("conn_handler: can't handle connection %v", err)
			c.conn.Close()
			return
		}
		// TODO: add len message in write data
		if _, err := c.conn.Write(res); err != nil {
			c.server.logger.Printf("conn_handler: can't write data in connection %v", err)
			c.conn.Close()
			return
		}
	}
}

func (c *ConnHandler) HandleClientCmd() (reply []byte, err error) {
	msgBuf := c.server.buffers.Get()
	defer c.server.buffers.Put(msgBuf)
	readData, err := c.conn.Read(msgBuf.Bytes())
	if err != nil {
		c.conn.Close()
		return nil, fmt.Errorf("conn_handler: can not read data from a connection: %v", err)
	}

	var cmd protocol.Command
	if readData > 0 && err == nil {
		if err := proto.Unmarshal(msgBuf.Bytes()[8:], &cmd); err != nil {	// fix problem with len msgBuf
			return nil, fmt.Errorf("conn_handler: can not Unmarshal read data: %v", err)
		}
	}

	result, err := c.executeCmd(&cmd)
	return c.encodeResponse(result, err)
}

func (c *ConnHandler) executeCmd(cmd *protocol.Command) (result []string, err error) {
	switch *cmd.Command {
	case protocol.CommandId_PING:
		result, err = c.Ping()
	case protocol.CommandId_HSET:
		result, err = c.HSet(cmd)
	case protocol.CommandId_HDEL:
		result, err = c.HDelete(cmd)
	case protocol.CommandId_HGET:
		result, err = c.HGet(cmd)
	case protocol.CommandId_ZADD:
		result, err = c.ZAdd(cmd)
	case protocol.CommandId_ZDELETE:
		result, err = c.ZDel(cmd)
	}
	return
}

func (c *ConnHandler) encodeResponse(values []string, err error) ([]byte, error) {
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}

	marshaledCmd, err := proto.Marshal(&protocol.Response{
		Values: values,
		Error:  &errMsg,
	})
	if err != nil {
		return nil, fmt.Errorf("conn_handler: can not Marshal reply data: %v", err)
	}

	return marshaledCmd, nil
}
