package cli

import "fmt"

func ObtainClient(host string, port int) {
	addr := fmt.Sprintf("%s:%d", host, port)
	return connect(addr, "tcp")
}
