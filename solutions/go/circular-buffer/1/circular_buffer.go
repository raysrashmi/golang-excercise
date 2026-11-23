package circularbuffer

import "errors"

// Define the Buffer type here.
type Buffer struct {
	data  []byte
	head  int // index of the oldest element
	count int // number of stored elements
}

func NewBuffer(size int) *Buffer {
	return &Buffer{data: make([]byte, size)}
}

func (b *Buffer) ReadByte() (byte, error) {
	if b.count == 0 {
		return 0, errors.New("buffer is empty")
	}
	c := b.data[b.head]
	b.head = (b.head + 1) % len(b.data)
	b.count--
	return c, nil
}

func (b *Buffer) WriteByte(c byte) error {
	if b.count == len(b.data) {
		return errors.New("buffer is full")
	}
	b.data[(b.head+b.count)%len(b.data)] = c
	b.count++
	return nil
}

func (b *Buffer) Overwrite(c byte) {
	if b.count == len(b.data) {
		b.data[b.head] = c
		b.head = (b.head + 1) % len(b.data)
		return
	}
	b.WriteByte(c)
}

func (b *Buffer) Reset() {
	b.head = 0
	b.count = 0
}
