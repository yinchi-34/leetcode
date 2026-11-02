package solution

import "sort"

// minMoves returns the minimum number of moves, where one move increments
// n-1 elements by 1. The answer is sum(nums) - n*min(nums).
func minMoves(nums []int) int {
	sum, mn := 0, nums[0]
	for _, x := range nums {
		sum += x
		if x < mn {
			mn = x
		}
	}
	return sum - len(nums)*mn
}

// minMovesBatched simulates the moves in batches: after sorting, the equal
// prefix nums[0..i-1] catches up with nums[i] by skipping nums[i] repeatedly.
// It does not modify nums.
func minMovesBatched(nums []int) int {
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
