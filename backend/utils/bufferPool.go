package utils

import (
	"bytes"
	"sync"
)

type BufferPool struct {
	pool sync.Pool
}

func (p *BufferPool) Get() *bytes.Buffer {
	buf := p.pool.Get()
	if buf == nil {
		return &bytes.Buffer{}
	}
	return buf.(*bytes.Buffer)
}

func (p *BufferPool) Put(buf *bytes.Buffer) {
	p.pool.Put(buf)
}
