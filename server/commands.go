package server

"github.com/DariaCrypto/Feather-KV-database/protocol"


// PING command.
func (c *ConnHandler) Ping() ([]string, error) {
	return []string{"Pong!"}, nil
}

/* HSET command. Set a value in a hashmap
   HSET <key> <val>*/
func (c *ConnHandler) HSet(cmd *protocol.Command) ([]string, error) {
}

/* HDEL command. Delete a value from a hashmap
   HDEL <key> <val>*/
   func (c *ConnHandler) HDelete(cmd *protocol.Command) ([]string, error) {
}

/* HGET command. Get a value from a hashmap
   HGET <key> <val>*/
   func (c *ConnHandler) HGet(cmd *protocol.Command) ([]string, error) {
}

/* ZADD command. Get a value from a sortedset
   ZADD <key> <val>*/
   func (c *ConnHandler) ZAdd(cmd *protocol.Command) ([]string, error) {
}

/* ZDEL command. Delete a value from a sortedset
   ZDEL <key> <val>*/
   func (c *ConnHandler) ZDel(cmd *protocol.Command) ([]string, error) {
}

/* ZMAX command. Get a max value from a sortedset
   ZMAX <key> <val>*/
   func (c *ConnHandler) ZMax(cmd *protocol.Command) ([]string, error) {
}

/* ZMIN command. Get a min value from a sortedset
   ZMIN <key> <val>*/
   func (c *ConnHandler) ZMin(cmd *protocol.Command) ([]string, error) {
}
