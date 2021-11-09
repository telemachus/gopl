package ch04

func RotateRight(nums []int, count int) []int {
	l := len(nums)
	if count < 1 || l < 2 {
		return nums
	}
	count %= l

	nums = append(nums[count:], nums[:count]...)
	return nums
}

func RotateLeft(nums []int, count int) []int {
	l := len(nums)
	if count < 1 || l < 2 {
		return nums
	}
	count %= l

	nums = append(nums[count:], nums[0:count]...)
	return nums
}

func AppendRotate(nums []int, n int) []int {
	if len(nums) <= 0 {
		return nums
	}

	r := len(nums) - n%len(nums)
	nums = append(nums[r:], nums[:r]...)
	return nums
}

func PopUnshiftRotate(nums []int, n int) []int {
	if len(nums) == 0 {
		return nums
	}
	n %= len(nums)
	if n == 0 {
		return nums
	}

	var popped int
	for i := 0; i < n; i++ {
		popped, nums = nums[len(nums)-1], nums[:len(nums)-1]
		nums = append([]int{popped}, nums...)
	}
	return nums
}

func ReverseRotate(nums []int, n int) {
	if len(nums) == 0 {
		return
	}
	n %= len(nums)
	if n == 0 {
		return
	}

	reverse(nums)
	reverse(nums[:n])
	reverse(nums[n:])
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
