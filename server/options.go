package server

import (
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

type Options struct {
	Address string
	LogFile string
}

type Option func(*Options)

func NewOptions(options ...Option) *Options {
	opt := &Options{}
	for _, option := range options {
		option(opt)
	}
	return opt
}

func WithAddress(addr string) Option {
	return func(o *Options) {
		o.Address = addr
	}
}

func WithLogFile(logname string) Option {
	return func(o *Options) {
		o.LogFile = logname
	}
}

func (o *Options) setLogger() *log.Logger {
	file, err := os.Create(o.LogFile)
	if err != nil {
		log.Printf("options: can not open/create log file: %v", err)
	}
	return log.New(file, "", log.LstdFlags|log.Lshortfile)
}

func (opt *Options) GetTCPAddress() *net.TCPAddr {
	addr := strings.Split(opt.GetAddress(), ":")
	port, _ := strconv.Atoi(addr[1])

	return &net.TCPAddr{
		IP:   net.ParseIP(addr[0]),
		Port: port,
	}
}

func (opt *Options) GetAddress() string {
	if opt.Address == "" {
		opt.Address = "127.0.0.1:6870"
	}

	return opt.Address
}
