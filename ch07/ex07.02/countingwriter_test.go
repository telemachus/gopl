package countingwriter_test

import (
	"io"
	"testing"

	"gopl/ch07/ex07.02/countingwriter"
)

func TestCountingWriter(t *testing.T) {
	cw := countingwriter.CountingWriter(io.Discard)
	bSliceOne := []byte("Hello, world!")
	bSliceTwo := []byte("Hello, 世界")

	t.Run("cw.Count() should be 0 initially", func(t *testing.T) {
		if cw.Count() != 0 {
			t.Errorf("expected: 0, actual: %d\n", cw.Count())
		}
	})

	a, _ := cw.Write(bSliceOne)
	t.Run("cw.Count() should be len(bSliceOne)", func(t *testing.T) {
		if int64(a) != cw.Count() {
			t.Errorf("expected: %d; actual: %d\n", a, cw.Count())
		}
	})

	b, _ := cw.Write(bSliceTwo)
	t.Run("cw.Count() should be len(bSliceOne)+len(bSliceTwo)", func(t *testing.T) {
		if int64(a+b) != cw.Count() {
			t.Errorf("expected: %d; actual: %d\n", a+b, cw.Count())
		}
	})

	t.Run("cw.Count() should be 0 after a reset", func(t *testing.T) {
		cw.Reset()
		if cw.Count() != 0 {
			t.Errorf("expected: 0; actual: %d\n", cw.Count())
		}
	})
}
