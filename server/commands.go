package server

import (
	"github.com/ddonskaya/feather/protocol"
)

// PING command.
func (c *ConnHandler) Ping() ([]string, error) {
	return []string{"Pong!"}, nil
}


//HSET command. Set a value in a hashmap
//HSET <key> <val>

func (c *ConnHandler) HSet(cmd *protocol.Command) ([]string, error) {
	c.server.hm.HSet(cmd.Args[0])
	return nil, nil
}


//HDEL command. Delete a value from a hashmap
//HDEL <key> <val>

func (c *ConnHandler) HDelete(cmd *protocol.Command) ([]string, error) {
	c.server.hm.HDelete(cmd.Args[0])
	return nil, nil
}


//HGET command. Get a value from a hashmap
//HGET <key> <val>

func (c *ConnHandler) HGet(cmd *protocol.Command) ([]string, error) {
	c.server.hm.HGet(cmd.Args[0])
	return nil, nil
}


//ZADD command. Get a value from a sortedset
//ZADD <key> <val>

func (c *ConnHandler) ZAdd(cmd *protocol.Command) ([]string, error) {
	c.server.ss.ZSet(cmd.Args[0])
	return nil, nil
}


//ZDEL command. Delete a value from a sortedset
//ZDEL <key> <val>

func (c *ConnHandler) ZDel(cmd *protocol.Command) ([]string, error) {
	c.server.ss.ZDelete(cmd.Args[0])
	return nil, nil
}


//ZMAX command. Get a max value from a sortedset
//ZMAX <key> <val>

func (c *ConnHandler) ZMax(cmd *protocol.Command) ([]string, error) {
}


//ZMIN command. Get a min value from a sortedset
//ZMIN <key> <val>

func (c *ConnHandler) ZMin(cmd *protocol.Command) ([]string, error) {
}
