package server

import (
	"net"

	"github.com/ddonskaya/feather/protocol"
	"github.com/ddonskaya/feather/utils"
	"google.golang.org/protobuf/proto"
)

type ConnHandler struct {
	conn      net.Conn
	server    *Server
	MsgHeader []byte
}

func newClient(conn net.Conn) *ConnHandler {
	return &ConnHandler{
		conn: conn,
	}
}

func (c *ConnHandler) HandleClientCmd() (reply []byte, err error) {
	cmd := new(protocol.Command)
	msgBuf := c.server.buffers.Get()
	defer c.server.buffers.Put(msgBuf)
	result, err := c.executeCmd(cmd)
	return c.encodeResponse(result, err)
}

func (c *ConnHandler) ProcessCmd(){
	for{
		res, err := c.HandleClientCmd()

		if err != nil {
			c.server.logger.Printf("conn_handler: can't handle connection %v", err)
			c.conn.Close()
			return
		}

		response := utils.UintToByteArray(uint64(len(res)))

		if _, err := c.conn.Write(append(response, res...)); err != nil {
			c.server.logger.Printf("conn_handler: can't write data in connection %v", err)
			c.conn.Close()
			return
		}
			
	}
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
	var errMsg string

	if err != nil {
		errMsg = err.Error()
	} else {
		errMsg = ""
	}

	return proto.Marshal(&protocol.Response{
		Values: values,
		Error:  &errMsg,
	})
}
