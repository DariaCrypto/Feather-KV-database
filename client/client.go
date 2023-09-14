package client

import (
	"errors"
	"fmt"
	"log"

	"github.com/ddonskaya/feather/protocol"
	"github.com/ddonskaya/feather/utils"
	"google.golang.org/protobuf/proto"
)

type FeatherClient struct {
	buffer    utils.BufferPool
	MsgHeader []byte
}

func NewFeatherClient() *FeatherClient {
	client := &FeatherClient{
		buffer:    *utils.NewBufferPool(),
		MsgHeader: make([]byte, utils.MSG_SIZE),
	}
	return client
}

func PerformCommand(c *FeatherClient, command *protocol.Command) (*protocol.Response, error) {
	cn, err := NewConnection()
	if err != nil {
		return nil, errors.New("client: сan not got connection failed")
	}

	marshaledCmd, err := proto.Marshal(command)
	if err != nil {
		return nil, errors.New("client: can not got marshaled data")
	}

	msgSize := utils.UintToByteArray(uint64(len(marshaledCmd)))
	msg := append(msgSize, marshaledCmd...)
	if _, err := cn.Write(msg); err != nil {
		log.Println("client: can not write data to connection", err)
		return nil, err
	}

	reply, err := c.getResponse(cn)
	return reply, err
}

func (c *FeatherClient) getResponse(conn *Connection) (*protocol.Response, error) {
	msgBuf := c.buffer.Get()
	defer c.buffer.Put(msgBuf)

	readData, err := conn.Read(msgBuf.Bytes())
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("client: can not read data from a connection: %v", err)
	}

	var response protocol.Response
	if readData > 0 && err == nil {	
		if err := proto.Unmarshal(msgBuf.Bytes()[:len(msgBuf.Bytes())-1], &response); err != nil { // fix problem with len msgBuf
			return nil, fmt.Errorf("conn_handler: can not Unmarshal read data: %v", err)
		}
	}
	return &response, nil
}
