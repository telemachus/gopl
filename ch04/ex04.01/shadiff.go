package ch04

func PopCountDifference(s1, s2 [32]uint8) int {
	popCountOne := sha256PopCount(s1)
	popCountTwo := sha256PopCount(s2)

	if popCountOne > popCountTwo {
		return popCountOne - popCountTwo
	}
	return popCountTwo - popCountOne
}

func sha256PopCount(shaSum [32]uint8) int {
	var n int
	for _, sum := range shaSum {
		n += popCount(sum)
	}

	return n
}

func popCount(x uint8) int {
	var n int
	for ; x > 0; x = x&(x-1) {
		n++
	}

	return n
}
