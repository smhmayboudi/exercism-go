package circular

import "errors"

type Buffer struct {
	data    []byte
	maxSize int
}

func NewBuffer(size int) *Buffer {
	return &Buffer{
		data:    []byte{},
		maxSize: size,
	}
}

func (b *Buffer) ReadByte() (byte, error) {
	if len(b.data) == 0 {
		return 0, errors.New("empty buffer")
	}

	val := b.data[0]
	b.data = b.data[1:]

	return val, nil
}

func (b *Buffer) WriteByte(c byte) error {
	if len(b.data) == b.maxSize {
		return errors.New("buffer full")
	}

	b.data = append(b.data, c)

	return nil
}

func (b *Buffer) Overwrite(c byte) {
	if len(b.data) == b.maxSize {
		b.data = b.data[1:]
	}
	b.data = append(b.data, c)
}

func (b *Buffer) Reset() {
	b.data = []byte{}
}
