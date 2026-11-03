package solution

import (
	"math/rand"
	"sort"
	"testing"

	lomuto "leetcode-go/Algorithms/order-statistics/0462-minimum-moves-to-equal-array-elements-ii/lomuto-quickselect"
	sorting "leetcode-go/Algorithms/order-statistics/0462-minimum-moves-to-equal-array-elements-ii/sorting"
	threeway "leetcode-go/Algorithms/order-statistics/0462-minimum-moves-to-equal-array-elements-ii/three-way-quickselect"
)

var solutions = []struct {
	name string
	fn   func([]int) int
}{
	{"sorting", sorting.MinMoves2},
	{"three_way_quickselect", threeway.MinMoves2},
	{"lomuto_quickselect", lomuto.MinMoves2},
}

var selectors = []struct {
	name string
	fn   func(a []int, l, r, k int) int
}{
	{"three_way_quickselect", threeway.QuickSelect},
	{"lomuto_quickselect", lomuto.QuickSelect},
}

func TestMinMoves2(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{name: "example_1", nums: []int{1, 2, 3}, expected: 2},
		{name: "example_2", nums: []int{1, 10, 2, 9}, expected: 16},
		{name: "single_element", nums: []int{7}, expected: 0},
		{name: "all_equal", nums: []int{5, 5, 5, 5, 5}, expected: 0},
		{name: "even_length", nums: []int{1, 1000}, expected: 999},
		{name: "outliers", nums: []int{1, 2, 3, 1000, 1001}, expected: 1998},
		{name: "negative_values", nums: []int{-5, 0, 3, -2}, expected: 10},
		{name: "many_duplicates", nums: []int{0, 1, 0, 1, 1, 0, 1}, expected: 3},
		{name: "extreme_values", nums: []int{-1000000000, 1000000000, 0}, expected: 2000000000},
	}

	for _, s := range solutions {
		for _, test := range tests {
			t.Run(s.name+"/"+test.name, func(t *testing.T) {
				result := s.fn(append([]int(nil), test.nums...))
				if result != test.expected {
					t.Errorf("%s(%v) = %d, want %d", s.name, test.nums, result, test.expected)
				}
			})
		}
	}
}

// TestSortingKeepsInput checks that the sorting solution works on a copy.
func TestSortingKeepsInput(t *testing.T) {
	nums := []int{3, 1, 2}
	sorting.MinMoves2(nums)
	if nums[0] != 3 || nums[1] != 1 || nums[2] != 2 {
		t.Fatalf("sorting.MinMoves2 modified its input: %v", nums)
	}
}

// TestQuickSelect checks every rank k against sorting for both partitions,
// on arrays with many duplicates, and checks that the result is a
// permutation of the input.
func TestQuickSelect(t *testing.T) {
	rng := rand.New(rand.NewSource(462))
	for iter := 0; iter < 2000; iter++ {
		n := 1 + rng.Intn(10)
		nums := make([]int, n)
		for i := range nums {
			nums[i] = rng.Intn(5) - 2
		}
		sorted := append([]int(nil), nums...)
		sort.Ints(sorted)
		for k := 0; k < n; k++ {
			for _, sel := range selectors {
				a := append([]int(nil), nums...)
				if got := sel.fn(a, 0, n-1, k); got != sorted[k] {
					t.Fatalf("%s(%v, k=%d) = %d, want %d", sel.name, nums, k, got, sorted[k])
				}
				sort.Ints(a)
				for i := range a {
					if a[i] != sorted[i] {
						t.Fatalf("%s(%v, k=%d) changed the multiset of elements", sel.name, nums, k)
					}
				}
			}
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// bruteForce tries every element as the common target; an optimal target
// can always be chosen among the elements.
func bruteForce(nums []int) int {
	best := -1
	for _, x := range nums {
		cost := 0
		for _, num := range nums {
			cost += abs(num - x)
		}
		if best < 0 || cost < best {
			best = cost
		}
	}
	return best
}

func TestAgainstBruteForce(t *testing.T) {
	rng := rand.New(rand.NewSource(4620))
	for iter := 0; iter < 1000; iter++ {
		n := 1 + rng.Intn(8)
		nums := make([]int, n)
		for i := range nums {
			nums[i] = rng.Intn(41) - 20
		}
		want := bruteForce(nums)
		for _, s := range solutions {
			if got := s.fn(append([]int(nil), nums...)); got != want {
				t.Fatalf("%s(%v) = %d, brute force = %d", s.name, nums, got, want)
			}
		}
	}
}

// TestAllEqualLarge guards against the quadratic behavior of two-way
// partitioning: with n identical values, three-way partitioning finishes
// in one pass. The Lomuto solution is left out on purpose: on the same
// input it makes about 3.7e9 comparisons and takes seconds.
func TestAllEqualLarge(t *testing.T) {
	nums := make([]int, 100000)
	if got := threeway.MinMoves2(nums); got != 0 {
		t.Fatalf("three-way MinMoves2(all zeros) = %d, want 0", got)
	}
}
