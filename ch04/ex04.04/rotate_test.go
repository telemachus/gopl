package rotate_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/telemachus/gopl/ch04/rotate"
)

type testCase struct {
	r     int
	orig  []int
	wantr []int
	wantl []int
}

func newTestCases() map[string]testCase {
	return map[string]testCase{
		"empty slice": {
			100,
			[]int{},
			[]int{},
			[]int{},
		},
		"one-item slice": {
			4,
			[]int{1},
			[]int{1},
			[]int{1},
		},
		"three-item slice": {
			1,
			[]int{1, 2, 3},
			[]int{3, 1, 2},
			[]int{2, 3, 1},
		},
		"eight-item slice": {
			1,
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
			[]int{8, 1, 2, 3, 4, 5, 6, 7},
			[]int{2, 3, 4, 5, 6, 7, 8, 1},
		},
		"rotation > length of slice": {
			7,
			[]int{1, 2, 3, 4, 5},
			[]int{4, 5, 1, 2, 3},
			[]int{3, 4, 5, 1, 2},
		},
		"negative rotation": {
			-1,
			[]int{1, 2, 3, 4, 5},
			[]int{2, 3, 4, 5, 1},
			[]int{5, 1, 2, 3, 4},
		},
		"negative rotation > length of slice": {
			-7,
			[]int{1, 2, 3, 4, 5},
			[]int{3, 4, 5, 1, 2},
			[]int{4, 5, 1, 2, 3},
		},
	}
}

func TestRight(t *testing.T) {
	t.Parallel()

	testCases := newTestCases()

	for msg, tc := range testCases {
		tc := tc

		t.Run(msg, func(t *testing.T) {
			t.Parallel()

			dupe := make([]int, len(tc.orig))
			copy(dupe, tc.orig)
			got := rotate.Right(dupe, tc.r)

			if !cmp.Equal(got, tc.wantr) {
				t.Errorf(
					"rotate.Right(%+v, %d) = %+v; want %+v\n",
					dupe,
					tc.r,
					got,
					tc.wantr,
				)
			}
		})
	}
}

func TestLeft(t *testing.T) {
	t.Parallel()

	testCases := newTestCases()

	for msg, tc := range testCases {
		tc := tc

		t.Run(msg, func(t *testing.T) {
			t.Parallel()

			dupe := make([]int, len(tc.orig))
			copy(dupe, tc.orig)
			got := rotate.Left(dupe, tc.r)

			if !cmp.Equal(got, tc.wantl) {
				t.Errorf(
					"rotate.Left(%+v, %d) = %+v; want %+v\n",
					dupe,
					tc.r,
					got,
					tc.wantl,
				)
			}
		})
	}
}

func TestReverseRight(t *testing.T) {
	t.Parallel()

	testCases := newTestCases()

	for msg, tc := range testCases {
		tc := tc

		t.Run(msg, func(t *testing.T) {
			t.Parallel()

			got := make([]int, len(tc.orig))
			copy(got, tc.orig)
			rotate.ReverseRight(got, tc.r)

			if !cmp.Equal(got, tc.wantr) {
				t.Errorf(
					"rotate.ReverseRight(%+v, %d) = %+v; want %+v\n",
					tc.orig,
					tc.r,
					got,
					tc.wantr,
				)
			}
		})
	}
}

func TestReverseLeft(t *testing.T) {
	t.Parallel()

	testCases := newTestCases()

	for msg, tc := range testCases {
		tc := tc

		t.Run(msg, func(t *testing.T) {
			t.Parallel()

			got := make([]int, len(tc.orig))
			copy(got, tc.orig)
			rotate.ReverseLeft(got, tc.r)

			if !cmp.Equal(got, tc.wantl) {
				t.Errorf(
					"rotate.ReverseLeft(%+v, %d) = %+v; want %+v\n",
					tc.orig,
					tc.r,
					got,
					tc.wantl,
				)
			}
		})
	}
}

func BenchmarkRightSmallSlice(b *testing.B) {
	nums := make([]int, 10)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		rotate.Right(nums, 5)
	}
}

func BenchmarkRightLargeSlice(b *testing.B) {
	nums := make([]int, 1000000)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		rotate.Right(nums, 500000)
	}
}

func BenchmarkReverseRightSmallSlice(b *testing.B) {
	nums := make([]int, 10)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		rotate.ReverseRight(nums, 5)
	}
}

func BenchmarkReverseRightLargeSlice(b *testing.B) {
	nums := make([]int, 1000000)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		rotate.ReverseRight(nums, 500000)
	}
}
