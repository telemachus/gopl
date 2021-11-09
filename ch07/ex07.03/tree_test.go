package main

import (
	"io"
	"testing"
)

func setup() *tree {
	t := add(nil, 1)
	for i := 2; i <= 100000; i++ {
		t = add(t, i)
	}
	return t
}

func BenchmarkPreorderRecursive(b *testing.B) {
	t := setup()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t.preorderPrint(io.Discard)
	}
}

func BenchmarkPostorderRecursive(b *testing.B) {
	t := setup()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t.postorderPrint(io.Discard)
	}
}

func BenchmarkInorderRecursive(b *testing.B) {
	t := setup()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t.inorderPrint(io.Discard)
	}
}

func BenchmarkMorris(b *testing.B) {
	t := setup()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		morrisPrint(io.Discard, t)
	}
}
