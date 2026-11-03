package solution

import "math/rand"

// partition is the two-way partition of CLRS §7.1 (Lomuto) with pivot
// x = a[r]: it returns q with a[l..q-1] <= x, a[q] = x, and a[q+1..r] > x.
func partition(a []int, l, r int) int {
	x := a[r]
	i := l - 1
	for j := l; j < r; j++ {
		if a[j] <= x {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[r] = a[r], a[i+1]
	return i + 1
}

func randomPartition(a []int, l, r int) int {
	p := rand.Intn(r-l+1) + l
	a[p], a[r] = a[r], a[p]
	return partition(a, l, r)
}

// QuickSelect returns the element that would be at index k if a[l..r] were
// sorted. It reorders a[l..r] in place. Every element equal to the pivot
// goes left, so many duplicates make it quadratic.
func QuickSelect(a []int, l, r, k int) int {
	q := randomPartition(a, l, r)
	if q == k {
		return a[q]
	}
	if q < k {
		return QuickSelect(a, q+1, r, k)
	}
	return QuickSelect(a, l, q-1, k)
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
