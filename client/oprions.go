package client

import (
	"time"
)

type Options struct {
	Network     string
	Address     string
	PollSize    int
	PoolTimeOut time.Duration
	IdleTimeout time.Duration
	MaxRetries int
}

type Option func(*Options)

func NewOptions(options ...Option) *Options{
    opt := &Options{}
    for _, option := range options {
        option(opt)
    }
    return opt
}

func WithNetwork(net string) Option{
	return func(o *Options) {
		o.Network = net
	}
}

func WithAddress(address string) Option{
	return func(o *Options) {
		o.Address = address
	}
}

func WithPollSize(size int) Option{
	return func(o *Options) {
		o.PollSize = size
	}
}

func WithPoolTimeOut(poolTimeOut time.Duration) Option{
	return func(o *Options) {
		o.PoolTimeOut = poolTimeOut
	}
}

func WithIdleTimeout(idleTimeOut time.Duration) Option{
	return func(o *Options) {
		o.IdleTimeout = idleTimeOut
	}
}

func WithMaxRetries(maxRetries int) Option{
	return func(o *Options) {
		o.MaxRetries = maxRetries
	}
}