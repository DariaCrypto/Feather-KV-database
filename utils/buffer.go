package utils

import (
	"bytes"
	"sync"
)

type BufferPool struct {
	pool *sync.Pool
}

func NewBuffer() *BufferPool {
	return &BufferPool{
		pool: &sync.Pool{
			New: func() interface{} {
				return bytes.Buffer{}
			},
		},
	}

}

func (bp *BufferPool) Put(any interface{}) {
	bp.pool.Put(any)
}

func (bp *BufferPool) Get() (any interface{}) {
	return bp.pool.Get()
}
