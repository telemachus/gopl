package ch04

func Reverse(ints *[3]int) {
	for i, j := 0, len(*ints)-1; i < j; i, j = i+1, j-1 {
		(*ints)[i], (*ints)[j] = (*ints)[j], (*ints)[i]
		// also works without the explicit pointers
		// ints[i], ints[j] = ints[j], ints[i]
	}
}
