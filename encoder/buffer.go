package encoder

import (
	"bytes"
	"errors"
)

/////////////////// WRITER PART ////////////////////////

type IWriter interface {
	WriteByte(byte) error
	Write([]byte) error
	GetContent() ([]byte, error)
}

type BufferWriter struct {
	buffer *bytes.Buffer
}

func (b *BufferWriter) WriteByte(data byte) error {
	return b.buffer.WriteByte(data)
}

func (b *BufferWriter) Write(datas []byte) error {
	_, err := b.buffer.Write(datas)
	return err
}

func (b *BufferWriter) GetContent() ([]byte, error) {
	raw := b.buffer.Bytes()
	b.buffer.Reset()
	result := make([]byte, len(raw))
	copy(result, raw)
	return result, nil
}

func NewBufferWriter() *BufferWriter {
	return &BufferWriter{
		buffer: bytes.NewBuffer(nil),
	}
}

type SliceWriter struct {
	slice []byte
}

func (b *SliceWriter) WriteByte(data byte) error {
	b.slice = append(b.slice, data)
	return nil
}

func (b *SliceWriter) Write(datas []byte) error {
	b.slice = append(b.slice, datas...)
	return nil
}

func (b *SliceWriter) GetContent() ([]byte, error) {
	result := make([]byte, len(b.slice))
	copy(result, b.slice)
	b.slice = b.slice[:0]
	return result, nil
}

func NewSliceWriter() *SliceWriter {
	return &SliceWriter{
		slice: []byte{},
	}
}

type _LinkedByteWriterNode struct {
	next *_LinkedByteWriterNode
	data []byte
}

/*
Make a linked list to store byte and convert them to byte array

I'm highly aware that there are data structure library for thing like this
Already checked some with 5K and 8K star but they didn't provide a simple and dynamic linked list implementation
So i have to make a simple wheel
*/
type LinkedByteWriter struct {
	start  *_LinkedByteWriterNode
	end    *_LinkedByteWriterNode
	result []byte
	length int
	size   int
}

func (l *LinkedByteWriter) WriteByte(data byte) error {
	return l.Write([]byte{data})
}

func (l *LinkedByteWriter) Write(data []byte) error {
	node := &_LinkedByteWriterNode{
		data: data,
	}
	if l.start == nil {
		l.start = node
		l.end = node
	} else {
		l.end.next = node
		l.end = node
	}
	l.length++
	l.size += len(data)

	//if new data is written then invalidate the cache
	l.result = nil
	return nil
}

func (l *LinkedByteWriter) GetContent() ([]byte, error) {
	if l.result == nil {
		if l.start == nil {
			l.result = []byte{}
		} else {
			l.result = l.toByte()
		}
	}
	return l.result, nil
}

func (l *LinkedByteWriter) toByte() []byte {
	iter := l.start
	//Should be able avoid reallocation due to resize
	data := make([]byte, 0, l.size)
	for iter != nil {
		data = append(data, iter.data...)
		iter = iter.next
	}
	return data
}

func (l *LinkedByteWriter) GetSize() int {
	return l.size
}

func NewLinkedByteWriter() *LinkedByteWriter {
	return &LinkedByteWriter{}
}

//////////////////// READER PART ///////////////////

type IReader interface {
	ReadByte() (byte, error)
	Read(int) ([]byte, error)
}

type SliceReader struct {
	buffer []byte
}

func (s *SliceReader) ReadByte() (byte, error) {
	if len(s.buffer) == 0 {
		return 0, errors.New("Not enough data in the buffer")
	}
	data := s.buffer[0]
	s.buffer = s.buffer[1:]
	return data, nil
}

func (s *SliceReader) Read(count int) ([]byte, error) {
	if len(s.buffer) < count {
		return nil, errors.New("Not enough data in the buffer")
	}
	data := s.buffer[:count]
	s.buffer = s.buffer[count:]
	return data, nil
}

func NewSliceReader(slice []byte) *SliceReader {
	return &SliceReader{
		buffer: slice,
	}
}
