package solution

import "sort"

// MinMoves simulates the moves in batches: after sorting, the equal prefix
// nums[0..i-1] catches up with nums[i] by skipping nums[i] repeatedly.
// It does not modify nums.
func MinMoves(nums []int) int {
	a := append([]int(nil), nums...)
	sort.Ints(a)

	res := 0
	for i := 1; i < len(a); i++ {
		cur := a[i] + res // nums[i] has never been skipped, so it gained res
		res += cur - a[i-1]
		a[i] = cur // the prefix now equals cur
	}
	return res
}
