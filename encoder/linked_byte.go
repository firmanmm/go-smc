package encoder

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
	start  *_LinkedByteNode
	end    *_LinkedByteNode
	result []byte
	length int
	size   int
}

func (l *LinkedByte) WriteByte(data byte) {
	l.Write([]byte{data})
}

func (l *LinkedByte) Write(data []byte) {
	node := &_LinkedByteNode{
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
}

func (l *LinkedByte) Merge(other *LinkedByte) {
	l.end.next = other.start
	l.end = other.end
	l.size += other.size
}

func (l *LinkedByte) GetResult() []byte {
	if l.result == nil {
		if l.start == nil {
			l.result = []byte{}
		} else {
			l.result = l.toByte()
		}
	}
	return l.result
}

func (l *LinkedByte) toByte() []byte {
	iter := l.start
	//Should be able avoid reallocation due to resize
	data := make([]byte, 0, l.size)
	for iter != nil {
		data = append(data, iter.data...)
		iter = iter.next
	}
	return data
}

func (l *LinkedByte) GetSize() int {
	return l.size
}

func NewLinkedByte() *LinkedByte {
	return &LinkedByte{}
}
