package utils

import (
	"bytes"
	"sync"
)

type BufferPool struct {
	pool *sync.Pool
}

func NewBufferPool() *BufferPool {
	return &BufferPool{
		pool: &sync.Pool{
			New: func() interface{} {
				return bytes.NewBuffer(make([]byte, 10))	// TODO: figure out with size
			},
		},
	}
}

func (bp *BufferPool) Put(buf *bytes.Buffer) {
	buf.Reset()
	bp.pool.Put(buf)
}

func (bp *BufferPool) Get() (*bytes.Buffer) {
	return bp.pool.Get().(*bytes.Buffer)
}
