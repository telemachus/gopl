package ch04_test

import (
	"crypto/sha256"
	"math/bits"
	"testing"

	"ch04"
)

func popCount(s [32]uint8) int {
	var n int
	for _, x := range s {
		n += bits.OnesCount8(x)
	}

	return n
}

func popCountDifference(s1, s2 [32]uint8) int {
	if s1 == s2 {
		return 0
	}

	s1Count := popCount(s1)
	s2Count := popCount(s2)

	if s1Count > s2Count {
		return s1Count - s2Count
	}
	return s2Count - s1Count
}

func TestSha256PopCountDifference(t *testing.T) {
	tests := map[string]struct {
		msg      string
		s1       [32]uint8
		s2       [32]uint8
		expected int
	}{
		"identical sha256s = 0 difference": {
			s1:       sha256.Sum256([]byte("x")),
			s2:       sha256.Sum256([]byte("x")),
			expected: 0,
		},
		"different sha256s != 0 difference": {
			s1:       sha256.Sum256([]byte("X")),
			s2:       sha256.Sum256([]byte("x")),
			expected: popCountDifference(sha256.Sum256([]byte("X")), sha256.Sum256([]byte("x"))),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := ch04.PopCountDifference(tc.s1, tc.s2)
			if tc.expected != actual {
				t.Errorf("ch04.PopCountDifference: expected %d but got %d\n", tc.expected, actual)
			}
		})
	}
}
