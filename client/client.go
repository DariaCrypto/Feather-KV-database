package client

import (
	"errors"
	"log"
	"github.com/ddonskaya/feather/protocol"
	"github.com/ddonskaya/feather/utils"
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

func NewFeatherClient() *FeatherClient {
	client := &FeatherClient{
		buffer: *utils.NewBuffer(),
		connPool: *NewConnectionPool(),
		MsgHeader: make([]byte, utils.MSG_SIZE),
	}
	return client
}

func PerformCommand(c *FeatherClient, command *protocol.Command) (*protocol.Response, error) {
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

	reply, err := c.getResponse(cn)
	return reply, err
}

func (c *FeatherClient) getResponse(conn Connection) (*protocol.Response, error){
	if _, err := utils.ReadData(conn, c.MsgHeader, utils.MSG_SIZE); err != nil {
		return nil, err
	}

	_, err := utils.ReadData(conn, c.MsgHeader, utils.MSG_SIZE)
	if err != nil {
		return nil, err
	}

	//Get id command
	idCmd := int(utils.ByteArrayToUint64(c.MsgHeader))
	msgBuf := c.buffer.Get().([]byte)
	defer c.buffer.Put(msgBuf)

	resp := &protocol.Response{}
	if err := proto.Unmarshal(msgBuf[:idCmd], resp); err != nil {
		return nil, err
	}

	return resp, nil
}
