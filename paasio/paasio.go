package paasio

import "io"

// Define readCounter and writeCounter types here.
type writeCounter struct {
	w io.Writer
}
type readCounter struct {
	r io.Reader
}

// For the return of the function NewReadWriteCounter, you must also define a type that satisfies the ReadWriteCounter interface.
type ReadWriteCount struct {
	WriteCounter
	ReadCounter
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{
		w: writer,
	}
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{
		r: reader,
	}
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	w := NewWriteCounter(readwriter)
	r := NewReadCounter(readwriter)
	return &ReadWriteCount{w, r}
}

func (rc *readCounter) Read(p []byte) (n int, err error) {
	return rc.r.Read(p)
}

func (rc *readCounter) ReadCount() (n int64, nops int) {
	panic("Please implement the ReadCount function")
}

func (wc *writeCounter) Write(p []byte) (n int, err error) {
	return wc.w.Write(p)
}

func (wc *writeCounter) WriteCount() (n int64, nops int) {
	panic("Please implement the WriteCount function")
}
