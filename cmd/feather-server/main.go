package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/ddonskaya/feather/server"
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
	Shutdown()
	opt := server.NewOptions(
		server.WithAddress("127.0.0.1:6870"),
		server.WithLogFile("LOG_FEATHER_SERVER.txt"),
	)

	go Shutdown()
	server.FeatherServer(opt)
}
