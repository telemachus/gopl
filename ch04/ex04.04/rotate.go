// Package rotate provides functions to rotate slices of any type.
package rotate

// Right rotates a slice of any type rightward by k steps and returns the
// rotated slice. The function properly handles any k. I.e., k can be zero,
// larger than the size of the slice, or negative.
func Right[T any](xs []T, k int) []T {
	// Avoid division by zero, and don't waste time on a singe-item slice.
	l := len(xs)
	if l < 2 {
		return xs
	}

	// Normalize rotations, both positive and negative, in case of overflow.
	// Go's % operator can lead to negative results, and that can cause
	// a panic from "slice bounds out of range." So, we use Euclidean mod
	// (see below) that always yields a positive remainder.
	r := mod((k + l), l)

	return append(xs[l-r:], xs[:l-r]...)
}

// Left rotates a slice of any type leftward by k steps and returns the rotated
// slice. The function properly handles any k. I.e., k can be zero, larger
// than the size of the slice, or negative.
func Left(xs []int, k int) []int {
	// Avoid division by zero, and don't waste time on a singe-item slice.
	l := len(xs)
	if l < 2 {
		return xs
	}

	// Normalize rotations, both positive and negative, in case of overflow.
	// Go's % operator can lead to negative results, and that can cause
	// a panic from "slice bounds out of range." So, we use Euclidean mod
	// (see below) that always yields a positive remainder.
	r := mod((k + l), l)

	return append(xs[r:], xs[:r]...)
}

// ReverseRight rotates a slice of any comparable type rightward by k steps.
// The rotation is done in place, so there is no return value. The function
// properly handles any k. I.e., k can be zero, larger than the size of the
// slice, or negative.
func ReverseRight[T comparable](xs []T, k int) {
	// Avoid division by zero, and don't waste time on a singe-item slice.
	l := len(xs)
	if l < 2 {
		return
	}

	// Normalize rotations, both positive and negative, in case of overflow.
	// Go's % operator can lead to negative results, and that can cause
	// a panic from "slice bounds out of range." So, we use Euclidean mod
	// (see below) that always yields a positive remainder.
	r := mod((k + l), l)
	if r == l {
		return
	}

	reverse(xs)
	reverse(xs[:r])
	reverse(xs[r:])
}

// ReverseLeft rotates a slice of any comparable type leftward by k steps.
// The rotation is done in place, so there is no return value. The function
// properly handles any k. I.e., k can be zero, larger than the size of the
// slice, or negative.
func ReverseLeft[T comparable](xs []T, k int) {
	// Avoid division by zero, and don't waste time on a singe-item slice.
	l := len(xs)
	if l < 2 {
		return
	}

	// Normalize rotations, both positive and negative, in case of overflow.
	// Go's % operator can lead to negative results, and that can cause
	// a panic from "slice bounds out of range." So, we use Euclidean mod
	// (see below) that always yields a positive remainder.
	r := mod((k + l), l)
	if r == l {
		return
	}

	reverse(xs[:r])
	reverse(xs[r:])
	reverse(xs)
}

// Return modulus with the sign of b rather than the sign of a, as Go normally does.
// See https://stackoverflow.com/a/59299881 and https://bit.ly/3YoeL1E.
// In our case, b is the length of the slice. Thus, the result of our use here
// is guaranteed to be positive.
func mod(a, b int) int {
	return (a%b + b) % b
}

func reverse[T comparable](xs []T) {
	for i, j := 0, len(xs)-1; i < j; i, j = i+1, j-1 {
		xs[i], xs[j] = xs[j], xs[i]
	}
}
