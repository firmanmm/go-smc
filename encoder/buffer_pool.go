package encoder

import (
	"bytes"
	"sync"
)

var bufferPool *sync.Pool

func _GetBuffer() *bytes.Buffer {
	buffer := bufferPool.Get().(*bytes.Buffer)
	buffer.Reset()
	return buffer
}

func _PutBuffer(buffer *bytes.Buffer) {
	bufferPool.Put(buffer)
}

var bufferTrackerPool *sync.Pool

func GetBufferTracker() *BufferTracker {
	buffer := bufferTrackerPool.Get().(*BufferTracker)
	buffer._Release()
	return buffer
}

func PutBufferTracker(bufferTracker *BufferTracker) {
	bufferTrackerPool.Put(bufferTracker)
}

type BufferTracker struct {
	backing []*bytes.Buffer
}

func (b *BufferTracker) Get() *bytes.Buffer {
	buffer := _GetBuffer()
	b.backing = append(b.backing, buffer)
	return buffer
}

func (b *BufferTracker) _Release() {
	for _, val := range b.backing {
		_PutBuffer(val)
	}
	b.backing = b.backing[:0]
}

func _NewBufferTracker() *BufferTracker {
	return &BufferTracker{
		backing: make([]*bytes.Buffer, 0, 2),
	}
}

func init() {
	bufferPool = &sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(nil)
		},
	}

	bufferTrackerPool = &sync.Pool{
		New: func() interface{} {
			return _NewBufferTracker()
		},
	}
}
