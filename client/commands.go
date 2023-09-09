package client

import (
	"fmt"
	"github.com/ddonskaya/feather/protocol"
)

//PING command
func Ping(c *FeatherClient) (*protocol.Response, error){
	cmdId := protocol.CommandId_PING

	cmd := &protocol.Command{
		Command: &cmdId,
	}
	
	responce, err := PerformCommand(c, cmd)
	if err != nil {
		return nil, fmt.Errorf("client: can't process PING command %v", err)
	}
	return responce, nil
}

//SET command
//SET <key> <value>
func Set(key, value string) {

}