package solution

import (
	"math/rand"
	"testing"
)

func TestMinMoves(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{name: "example_1", nums: []int{2, 1, 3}, expected: 3},
		{name: "example_2", nums: []int{4, 4, 5}, expected: 2},
		{name: "single_element", nums: []int{7}, expected: 0},
		{name: "all_equal", nums: []int{3, 3, 3}, expected: 0},
		{name: "max_first", nums: []int{9, 1, 5}, expected: 12},
		{name: "repeated_maximum", nums: []int{5, 2, 5}, expected: 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := MinMoves(test.nums); got != test.expected {
				t.Errorf("MinMoves(%v) = %d, want %d", test.nums, got, test.expected)
			}
		})
	}
}

// bruteForce tries every target value from max to max+10 and keeps the cheapest.
func bruteForce(nums []int) int {
	mx := nums[0]
	for _, v := range nums {
		if v > mx {
			mx = v
		}
	}
	best := -1
	for target := mx; target <= mx+10; target++ {
		cost := 0
		for _, v := range nums {
			cost += target - v
		}
		if best < 0 || cost < best {
			best = cost
		}
	}
	return best
}

func TestAgainstBruteForce(t *testing.T) {
	rng := rand.New(rand.NewSource(3736))
	for iter := 0; iter < 2000; iter++ {
		n := 1 + rng.Intn(10)
		nums := make([]int, n)
		for i := range nums {
			nums[i] = 1 + rng.Intn(100)
		}
		if got, want := MinMoves(nums), bruteForce(nums); got != want {
			t.Fatalf("MinMoves(%v) = %d, want %d", nums, got, want)
		}
	}
}
