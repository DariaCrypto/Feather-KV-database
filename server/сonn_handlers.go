package server

import (
	"errors"
	"feather-kv/protocol"
	"feather-kv/utils"
	"net"
	"github.com/DariaCrypto/Feather-KV-database/feather"

	"google.golang.org/protobuf/proto"
)

type ConnHandler struct {
	conn      net.Conn
	server    *Server
	MsgHeader []byte
	hm *feather.HashMapCollection
	ss *feather.SortedSetCollection
}

func newClient(conn net.Conn) *ConnHandler {
	return &ConnHandler{
		conn: conn,
	}
}

func (c *ConnHandler) Exec() (reply []byte, err error) {

	cmd := new(protocol.Command)

	if read, err := utils.ReadData(c.conn, c.MsgHeader, utils.MSG_SIZE); err != nil {
		return nil, errors.New("")
	}

}

func (c *ConnHandler) encodeReply(values []string, err error) ([]byte, error) {
	var errMsg string

	if err != nil {
		errMsg = err.Error()
	} else {
		errMsg = ""
	}

	return proto.Marshal(&protocol.Reply{
		Values: values,
		Error:  &errMsg,
	})
}
