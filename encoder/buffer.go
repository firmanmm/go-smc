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
	return b.buffer.Bytes(), nil
}

func NewBufferWriter() *BufferWriter {
	return &BufferWriter{
		buffer: bytes.NewBuffer(nil),
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
