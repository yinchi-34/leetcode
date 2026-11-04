package solution

import (
	"math/rand"
	"testing"

	batched "leetcode-go/Algorithms/simulation/0453-minimum-moves-to-equal-array-elements/batched-simulation"
	formula "leetcode-go/Algorithms/simulation/0453-minimum-moves-to-equal-array-elements/formula"
)

var solutions = []struct {
	name string
	fn   func([]int) int
}{
	{"batched_simulation", batched.MinMoves},
	{"formula", formula.MinMoves},
}

func TestMinMoves(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{name: "example_1", nums: []int{1, 2, 3}, expected: 3},
		{name: "example_2", nums: []int{1, 1, 1}, expected: 0},
		{name: "single_element", nums: []int{7}, expected: 0},
		{name: "unsorted", nums: []int{5, 1, 7, 2}, expected: 11},
		{name: "negative_values", nums: []int{-3, -1, 2}, expected: 7},
		{name: "repeated_minimum", nums: []int{2, 2, 5}, expected: 3},
		{name: "large_gap", nums: []int{-1000000000, 1000000000}, expected: 2000000000},
	}

	for _, s := range solutions {
		for _, test := range tests {
			t.Run(s.name+"/"+test.name, func(t *testing.T) {
				input := append([]int(nil), test.nums...)
				result := s.fn(input)
				if result != test.expected {
					t.Errorf("%s(%v) = %d, want %d", s.name, test.nums, result, test.expected)
				}
				for i := range input {
					if input[i] != test.nums[i] {
						t.Fatalf("%s modified its input: got %v, want %v", s.name, input, test.nums)
					}
				}
			})
		}
	}
}

// simulate performs the moves one at a time, always skipping a maximum.
func simulate(nums []int) int {
	a := append([]int(nil), nums...)
	moves := 0
	for {
		mx, mn := 0, 0
		for i := range a {
			if a[i] > a[mx] {
				mx = i
			}
			if a[i] < a[mn] {
				mn = i
			}
		}
		if a[mx] == a[mn] {
			return moves
		}
		for i := range a {
			if i != mx {
				a[i]++
			}
		}
		moves++
	}
}

func TestAgainstSimulation(t *testing.T) {
	rng := rand.New(rand.NewSource(453))
	for iter := 0; iter < 500; iter++ {
		n := 1 + rng.Intn(6)
		nums := make([]int, n)
		for i := range nums {
			nums[i] = rng.Intn(21) - 10
		}
		want := simulate(nums)
		for _, s := range solutions {
			if got := s.fn(nums); got != want {
				t.Fatalf("%s(%v) = %d, simulation = %d", s.name, nums, got, want)
			}
		}
	}
}
