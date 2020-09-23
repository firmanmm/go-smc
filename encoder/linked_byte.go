package encoder

import (
	"bytes"
)

type _LinkedByteNode struct {
	next *_LinkedByteNode
	data []byte
}

/*
Make a linked list to store byte and convert them to byte array

I'm highly aware that there are data structure library for thing like this
Already checked some with 5K and 8K star but they didn't provide a simple and dynamic linked list implementation
So i have to make a simple wheel
*/
type LinkedByte struct {
	buffer *bytes.Buffer
	result []byte
}

func (l *LinkedByte) WriteByte(data byte) {
	l.buffer.WriteByte(data)
}

func (l *LinkedByte) Write(data []byte) {
	l.buffer.Write(data)

	//if new data is written then invalidate the cache
	l.result = nil
}

func (l *LinkedByte) Merge(other *LinkedByte) {
	other.buffer.WriteTo(l.buffer)
}

func (l *LinkedByte) GetResult() []byte {
	if l.result == nil {
		l.result = l.toByte()
	}
	return l.result
}

func (l *LinkedByte) toByte() []byte {
	return l.buffer.Bytes()
}

func (l *LinkedByte) GetSize() int {
	return l.buffer.Len()
}

func NewLinkedByte() *LinkedByte {
	return &LinkedByte{
		buffer: bytes.NewBuffer(nil),
	}
}
