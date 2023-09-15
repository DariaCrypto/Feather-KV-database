package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/ddonskaya/feather/server"
	"github.com/ddonskaya/feather/utils"
)

func Shutdown() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	go func() {
		<-c
		os.Exit(0)
	}()
}

func main() {
	opt := server.NewOptions(
		server.WithAddress(utils.SERVER),
		server.WithLogFile("LOG_FEATHER_SERVER.txt"),
	)

	go Shutdown()
	server.FeatherServer(opt)
}
