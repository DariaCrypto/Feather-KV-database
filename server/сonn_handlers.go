package server

import (
	"errors"
	"net"

	"github.com/ddonskaya/feather/collections"
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

	if read, err := utils.ReadData(c.conn, c.MsgHeader, utils.MSG_SIZE); err != nil {
		return nil, errors.New("")
	}

	idCmd := int(utils.ByteArrayToUint64(c.MsgHeader))
	msgBuf := c.server.buffers.Get()
	defer c.server.buffers.Put()

	if read, err := utils.ReadData(c.conn, c.MsgHeader, utils.MSG_SIZE); err != nil {
			return nil, errors.New("")
	}

	result, err := c.executeCmd(cmd)
	return c.encodeResponse(result, err)
}

func (c *ConnHandler) executeCmd(cmd *protocol.Command) (result []string, err error) {
	switch *cmd.CommandId {
	case protocol.CommandI_PING:
		result, err = c.Ping()
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

