package solution

import "sort"

// MinMoves2 finds a median by sorting a copy of nums and returns the total
// distance of all elements to it. It does not modify nums.
func MinMoves2(nums []int) int {
	a := append([]int(nil), nums...)
	sort.Ints(a)
	x := a[len(a)/2]
	res := 0
	for _, num := range a {
		res += abs(num - x)
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
