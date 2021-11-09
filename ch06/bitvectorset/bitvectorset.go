package bitvectorset

import (
	"bytes"
	"fmt"
	"math/bits"
)

const uintBitSize = 32 << (^uint(0) >> 63)

type IntSet struct {
	words []uint
}

// Len reports the number of items in the set.
func (s *IntSet) Len() int {
	var n int
	for _, word := range s.words {
		n += bits.OnesCount(word)
	}
	return n
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/uintBitSize, uint(x%uintBitSize)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// AddAll is a variadic version of Add.
func (s *IntSet) AddAll(xs ...int) {
	for _, x := range xs {
		word, bit := x/uintBitSize, uint(x%uintBitSize)
		for word >= len(s.words) {
			s.words = append(s.words, 0)
		}
		s.words[word] |= 1 << bit
	}
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/uintBitSize, uint(x%uintBitSize)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Remove removes the non-negative value x from the set.
func (s *IntSet) Remove(x int) {
	word, bit := x/uintBitSize, uint(x%uintBitSize)
	// for word >= len(s.words) {
	// 	s.words = append(s.words, 0)
	// }
	s.words[word] &^= 1 << bit
}

// Clear removes all elements from the set.
func (s *IntSet) Clear() {
	s.words = nil
}

// Copy creates and returns a copy of the set.
func (s *IntSet) Copy() *IntSet {
	newSet := &IntSet{}
	newSet.words = make([]uint, len(s.words))
	copy(newSet.words, s.words)
	return newSet
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWith sets s to the intersection of s and t.
func (s *IntSet) IntersectWith(t *IntSet) {
	tLength := len(t.words)
	for i := 0; i < len(s.words); i++ {
		if i < tLength {
			s.words[i] &= t.words[i]
		} else {
			s.words[i] &= 0
		}
	}
}

// DifferenceWith sets s to the difference between s and t.
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words[:len(s.words)] {
		s.words[i] &^= tword
	}
}

// SymmetricDifferenceWith sets s to the symmetric difference of s and t.
func (s *IntSet) SymmetricDifferenceWith(t *IntSet) {
	// If s.words has fewer items than t.words, we do this in two steps.
	// First, use XOR to remove items in t.words[:sLen] from s.words.
	// Second, add everything from t.words[sLen:] to the shorter set.
	//
	// Otherwise, we only need to remove items in t.words from s.words, and
	// there are no complications about where to start or end iterating.
	sLen, tLen := len(s.words), len(t.words)
	if sLen < tLen {
		for i, tword := range t.words[:sLen] {
			s.words[i] ^= tword
		}
		for _, tword := range t.words[sLen:] {
			s.words = append(s.words, tword)
		}
	} else {
		for i, tword := range t.words {
			s.words[i] ^= tword
		}
	}
}

// Same reports whether s is the same set as t.
func (s *IntSet) Same(t *IntSet) bool {
	if len(s.words) != len(t.words) {
		return false
	}
	for i, word := range s.words {
		if word != t.words[i] {
			return false
		}
	}
	return true
}

func (s *IntSet) Elems() []int {
	es := make([]int, 0, len(s.words))
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < uintBitSize; j++ {
			if word&(1<<uint(j)) != 0 {
				es = append(es, uintBitSize*i+j)
			}
		}
	}
	return es
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < uintBitSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(',')
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", uintBitSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
