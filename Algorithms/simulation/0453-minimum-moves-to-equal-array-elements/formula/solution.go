package solution

// MinMoves returns the minimum number of moves, where one move increments
// n-1 elements by 1. The answer is sum(nums) - n*min(nums).
func MinMoves(nums []int) int {
	sum, mn := 0, nums[0]
	for _, x := range nums {
		sum += x
		if x < mn {
			mn = x
		}
	}
	return sum - len(nums)*mn
}
