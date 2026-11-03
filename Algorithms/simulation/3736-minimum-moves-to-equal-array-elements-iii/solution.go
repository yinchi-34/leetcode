package solution

// MinMoves returns the minimum number of moves to make all elements equal,
// where one move increases a single element by 1.
// The final value is at least max(nums), and choosing exactly max(nums)
// costs sum(max - nums[i]) = n*max - sum.
func MinMoves(nums []int) int {
	sum, mx := 0, nums[0]
	for _, num := range nums {
		sum += num
		if num > mx {
			mx = num
		}
	}
	return len(nums)*mx - sum
}
