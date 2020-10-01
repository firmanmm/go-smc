package encoder

import (
	"bytes"
	"errors"
)

/////////////////// WRITER PART ////////////////////////

type IWriter interface {
	WriteString(string) error
	WriteByte(byte) error
	Write([]byte) error
	GetContent() ([]byte, error)
}

type BufferWriter struct {
	buffer *bytes.Buffer
}

func (b *BufferWriter) WriteString(data string) error {
	_, err := b.buffer.WriteString(data)
	return err
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
	data [][]byte
}

func New_LinkedByteWriterNode(size int) *_LinkedByteWriterNode {
	return &_LinkedByteWriterNode{
		data: make([][]byte, size),
	}
}

/*
Make a linked list to store byte and convert them to byte array

I'm highly aware that there are data structure library for thing like this
Already checked some with 5K and 8K star but they didn't provide a simple and dynamic linked list implementation
So i have to make a simple wheel
//DEV NOTE
Sometime, out smarting the system just isn't worth it, This particular writer is slow, i mean, really slow.
Sticking with bytes buffer or even slice append is faster
*/
type LinkedByteWriter struct {
	start *_LinkedByteWriterNode
	end   *_LinkedByteWriterNode
	size  int
	idx   int
}

func (l *LinkedByteWriter) WriteString(data string) error {
	return l.Write([]byte(data))
}

func (l *LinkedByteWriter) WriteByte(data byte) error {
	return l.Write([]byte{data})
}

func (l *LinkedByteWriter) Write(data []byte) error {
	if l.idx == len(l.end.data) {
		node := New_LinkedByteWriterNode(l.idx * 2)
		l.idx = 0
		l.end.next = node
		l.end = node
	}
	l.end.data[l.idx] = data
	l.idx++
	l.size += len(data)
	return nil
}

func (l *LinkedByteWriter) GetContent() ([]byte, error) {
	res := l.toByte()
	l.end = l.start
	l.idx = 0
	return res, nil
}

func (l *LinkedByteWriter) toByte() []byte {
	//Should be able avoid reallocation due to resize
	data := make([]byte, 0, l.size)
	iter := l.start
	for iter != nil {
		limit := len(iter.data)
		if iter.next == nil {
			limit = l.idx
		}

		for i := 0; i < limit; i++ {
			data = append(data, iter.data[i]...)
		}
		iter = iter.next
	}
	return data
}

func (l *LinkedByteWriter) GetSize() int {
	return l.size
}

func NewLinkedByteWriter() *LinkedByteWriter {
	node := New_LinkedByteWriterNode(2)
	return &LinkedByteWriter{
		start: node,
		end:   node,
	}
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
