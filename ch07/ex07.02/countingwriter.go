package countingwriter

import "io"

type writeCounter struct {
	w     io.Writer
	count int64
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	n, err := wc.w.Write(p)
	wc.count += int64(n)
	return n, err
}

func (wc *writeCounter) Count() int64 {
	return wc.count
}

func (wc *writeCounter) Reset() {
	wc.count = 0
}

func CountingWriter(w io.Writer) *writeCounter {
	return &writeCounter{w, 0}
}
