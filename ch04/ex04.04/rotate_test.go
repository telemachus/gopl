package ch04_test

import (
	"ch04"
	"testing"
)

func compareIntSlices(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, n := range s1 {
		if n != s2[i] {
			return false
		}
	}
	return true
}

func TestReverseRotate(t *testing.T) {
	tests := map[string]struct {
		rotations int
		nums      []int
		expected  []int
	}{
		"empty slice": {rotations: 5, nums: []int{}, expected: []int{}},
		"odd slice even rotations": {rotations: 2, nums: []int{1,2,3}, expected: []int{2,3,1}},
		"even slice even rotations": {rotations: 2, nums: []int{1,2,3,4}, expected: []int{3,4,1,2}},
		"odd slice odd rotations": {rotations: 3, nums: []int{1,2,3}, expected: []int{1,2,3}},
		"even slice odd rotations": {rotations: 3, nums: []int{1,2,3,4}, expected: []int{2,3,4,1}},
		"rotations > len(slice)": {rotations: 5, nums: []int{1,2,3}, expected: []int{2,3,1}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			ch04.ReverseRotate(tc.nums, tc.rotations)
			if !compareIntSlices(tc.expected, tc.nums) {
				t.Errorf("expected %#v; actual %#v", tc.expected, tc.nums)
			}
		})
	}
}

func BenchmarkReverseRotateSmallSlice(b *testing.B) {
	nums := []int{1,2,3,4,5,6,7,8,9,10}
	for i := 0; i < b.N; i++ {
		ch04.ReverseRotate(nums, 5)
	}
}

func BenchmarkReverseRotateLargeSlice(b *testing.B) {
	nums := make([]int, 0, 100000)
	for i := 1; i <= 100000; i++ {
		nums = append(nums, i)
	}
	for i := 0; i < b.N; i++ {
		ch04.ReverseRotate(nums, 100)
	}
}

func TestReversePopUnshiftRotate(t *testing.T) {
	tests := map[string]struct {
		rotations int
		nums      []int
		expected  []int
	}{
		"empty slice": {rotations: 5, nums: []int{}, expected: []int{}},
		"odd slice even rotations": {rotations: 2, nums: []int{1,2,3}, expected: []int{2,3,1}},
		"even slice even rotations": {rotations: 2, nums: []int{1,2,3,4}, expected: []int{3,4,1,2}},
		"odd slice odd rotations": {rotations: 3, nums: []int{1,2,3}, expected: []int{1,2,3}},
		"even slice odd rotations": {rotations: 3, nums: []int{1,2,3,4}, expected: []int{2,3,4,1}},
		"rotations > len(slice)": {rotations: 5, nums: []int{1,2,3}, expected: []int{2,3,1}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tc.nums = ch04.PopUnshiftRotate(tc.nums, tc.rotations)
			if !compareIntSlices(tc.expected, tc.nums) {
				t.Errorf("expected %#v; actual %#v", tc.expected, tc.nums)
			}
		})
	}
}

func BenchmarkPopUnshiftRotateSmallSlice(b *testing.B) {
	nums := []int{1,2,3,4,5,6,7,8,9,10}
	for i := 0; i < b.N; i++ {
		ch04.PopUnshiftRotate(nums, 5)
	}
}

func BenchmarkPopUnshiftRotateLargeSlice(b *testing.B) {
	nums := make([]int, 0, 100000)
	for i := 1; i <= 100000; i++ {
		nums = append(nums, i)
	}
	for i := 0; i < b.N; i++ {
		ch04.PopUnshiftRotate(nums, 100)
	}
}

func TestAppendRotate(t *testing.T) {
	tests := map[string]struct {
		rotations int
		nums      []int
		expected  []int
	}{
		"empty slice": {rotations: 5, nums: []int{}, expected: []int{}},
		"odd slice even rotations": {rotations: 2, nums: []int{1,2,3}, expected: []int{2,3,1}},
		"even slice even rotations": {rotations: 2, nums: []int{1,2,3,4}, expected: []int{3,4,1,2}},
		"odd slice odd rotations": {rotations: 3, nums: []int{1,2,3}, expected: []int{1,2,3}},
		"even slice odd rotations": {rotations: 3, nums: []int{1,2,3,4}, expected: []int{2,3,4,1}},
		"rotations > len(slice)": {rotations: 5, nums: []int{1,2,3}, expected: []int{2,3,1}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tc.nums = ch04.AppendRotate(tc.nums, tc.rotations)
			if !compareIntSlices(tc.expected, tc.nums) {
				t.Errorf("expected %#v; actual %#v", tc.expected, tc.nums)
			}
		})
	}
}

func BenchmarkAppendRotateSmallSlice(b *testing.B) {
	nums := []int{1,2,3,4,5,6,7,8,9,10}
	for i := 0; i < b.N; i++ {
		ch04.AppendRotate(nums, 5)
	}
}

func BenchmarkAppendRotateLargeSlice(b *testing.B) {
	nums := make([]int, 0, 100000)
	for i := 1; i <= 100000; i++ {
		nums = append(nums, i)
	}
	for i := 0; i < b.N; i++ {
		ch04.AppendRotate(nums, 100)
	}
}

func TestRotateRight(t *testing.T) {
	tests := map[string]struct {
		rotations int
		nums      []int
		expected  []int
	}{
		"empty slice": {rotations: 5, nums: []int{}, expected: []int{}},
		"odd slice even rotations": {rotations: 2, nums: []int{1,2,3}, expected: []int{3,1,2}},
		"even slice even rotations": {rotations: 2, nums: []int{1,2,3,4}, expected: []int{3,4,1,2}},
		"odd slice odd rotations": {rotations: 3, nums: []int{1,2,3}, expected: []int{1,2,3}},
		"even slice odd rotations": {rotations: 3, nums: []int{1,2,3,4}, expected: []int{4,1,2,3}},
		"rotations > len(slice)": {rotations: 5, nums: []int{1,2,3}, expected: []int{3,1,2}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tc.nums = ch04.RotateRight(tc.nums, tc.rotations)
			if !compareIntSlices(tc.expected, tc.nums) {
				t.Errorf("expected %#v; actual %#v", tc.expected, tc.nums)
			}
		})
	}
}

func BenchmarkRotateRightSmallSlice(b *testing.B) {
	nums := []int{1,2,3,4,5,6,7,8,9,10}
	for i := 0; i < b.N; i++ {
		ch04.RotateRight(nums, 5)
	}
}

func BenchmarkRotateRightLargeSlice(b *testing.B) {
	nums := make([]int, 0, 100000)
	for i := 1; i <= 100000; i++ {
		nums = append(nums, i)
	}
	for i := 0; i < b.N; i++ {
		ch04.RotateRight(nums, 100)
	}
}

func TestRotateLeft(t *testing.T) {
	tests := map[string]struct {
		rotations int
		nums      []int
		expected  []int
	}{
		"empty slice": {rotations: 5, nums: []int{}, expected: []int{}},
		"odd slice even rotations": {rotations: 2, nums: []int{1,2,3}, expected: []int{3,1,2}},
		"even slice even rotations": {rotations: 2, nums: []int{1,2,3,4}, expected: []int{3,4,1,2}},
		"odd slice odd rotations": {rotations: 3, nums: []int{1,2,3}, expected: []int{1,2,3}},
		"even slice odd rotations": {rotations: 3, nums: []int{1,2,3,4}, expected: []int{4,1,2,3}},
		"rotations > len(slice)": {rotations: 5, nums: []int{1,2,3}, expected: []int{3,1,2}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tc.nums = ch04.RotateLeft(tc.nums, tc.rotations)
			if !compareIntSlices(tc.expected, tc.nums) {
				t.Errorf("expected %#v; actual %#v", tc.expected, tc.nums)
			}
		})
	}
}

func BenchmarkRotateLeftSmallSlice(b *testing.B) {
	nums := []int{1,2,3,4,5,6,7,8,9,10}
	for i := 0; i < b.N; i++ {
		ch04.RotateLeft(nums, 5)
	}
}

func BenchmarkRotateLeftLargeSlice(b *testing.B) {
	nums := make([]int, 0, 100000)
	for i := 1; i <= 100000; i++ {
		nums = append(nums, i)
	}
	for i := 0; i < b.N; i++ {
		ch04.RotateLeft(nums, 100)
	}
}
