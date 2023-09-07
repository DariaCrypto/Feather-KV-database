package client

import (
	"context"
	"errors"
	"log"
	"time"
	"net"
	"github.com/DariaCrypto/Feather-KV-database/utils"
	"github.com/DariaCrypto/Feather-KV-database/protocol"
	"google.golang.org/protobuf/proto"
)

const (
	serverAddress = "127.0.0.1:8080"
	numClients    = 10
)

type FeatherClient struct {
	buffer utils.BufferPool
	connPool  ConnPool
	MsgHeader []byte
}

func New() *FeatherClient {
	client := &FeatherClient{
		buffer: *utils.NewBuffer(),
		connPool: *NewConnectionPool(),
	}
	return client
}

func (c *FeatherClient) process(command *protocol.Command) (*protocol.Reply, error) {
	//Get a connection from poolConnection
	cn, err := c.connPool.Get()
	defer c.connPool.Put(cn)
	if err != nil {
		return nil, errors.New("client: Can't got connection from connPool")
	}

	marshaledCmd, err := proto.Marshal(command)
	if err != nil {
		return nil, errors.New("client: can not got marshaled data")
	}

	msgSize := utils.UintToByteArray(uint64(len(marshaledCmd)))
	msg := append(msgSize, marshaledCmd...)

	if _, err := cn.Write(msg); err != nil {
		log.Println("client: ", err)
		return nil, err
	}

	reply, err := c.getReply(&cn)
	return reply, err
}

func (c *FeatherClient) getReply(conn *net.Conn) (*protocol.Reply, error){
	if read, err := utils.ReadData(conn, c.MsgHeader, utils.MSG_SIZE); err != nil {
		return nil, err
	}

	//Get id command
	idCmd := int(utils.ByteArrayToUint64(c.MsgHeader))
	msgBuf := c.buffer.Get().([]byte)
	defer c.buffer.Put(msgBuf)
	read, err := utils.ReadData(conn, c.MsgHeader, utils.MSG_SIZE)
	if err != nil {
		return nil, err
	}

	reply := &protocol.Reply{}
	if err := proto.Unmarshal(msgBuf[:idCmd], reply); err != nil {
		return nil, err
	}

	return reply, nil
}
