package solution

import "math/rand"

// partition3 rearranges a[l..r] around the pivot x = a[r] into three parts:
// a[l..lt-1] < x, a[lt..gt] == x, a[gt+1..r] > x, and returns lt and gt.
func partition3(a []int, l, r int) (int, int) {
	x := a[r]
	lt, i, gt := l, l, r
	for i <= gt {
		switch {
		case a[i] < x:
			a[lt], a[i] = a[i], a[lt]
			lt++
			i++
		case a[i] > x:
			a[i], a[gt] = a[gt], a[i]
			gt-- // the element swapped in from gt has not been examined yet
		default:
			i++
		}
	}
	return lt, gt
}

// randomPartition3 moves a uniformly random element of a[l..r] to a[r]
// and partitions around it.
func randomPartition3(a []int, l, r int) (int, int) {
	p := rand.Intn(r-l+1) + l
	a[p], a[r] = a[r], a[p]
	return partition3(a, l, r)
}

// QuickSelect returns the element that would be at index k if a[l..r] were
// sorted. It reorders a[l..r] in place.
func QuickSelect(a []int, l, r, k int) int {
	lt, gt := randomPartition3(a, l, r)
	if k < lt {
		return QuickSelect(a, l, lt-1, k)
	}
	if k > gt {
		return QuickSelect(a, gt+1, r, k)
	}
	return a[k]
}

// MinMoves2 moves every element to a median found by QuickSelect.
// It reorders nums in place.
func MinMoves2(nums []int) int {
	x := QuickSelect(nums, 0, len(nums)-1, len(nums)/2)
	res := 0
	for _, num := range nums {
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
